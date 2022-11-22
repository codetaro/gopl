package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

type zFunc func(x, y float64) float64

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// parse zFunc
		var z zFunc
		switch request.URL.Query().Get("zFunc") {
		case "peak":
			z = peak
		case "eggbox":
			z = eggbox
		}

		writer.Header().Set("Content-Type", "image/svg+xml")
		surface(writer, z)
	})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func surface(w io.Writer, z zFunc) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width:0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, err := corner(i+1, j, z)
			if err != nil {
				continue
			}
			bx, by, err := corner(i, j, z)
			if err != nil {
				continue
			}
			cx, cy, err := corner(i, j+1, z)
			if err != nil {
				continue
			}
			dx, dy, err := corner(i+1, j+1, z)
			if err != nil {
				continue
			}
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(w, "</svg>")
}

func corner(i, j int, f zFunc) (float64, float64, error) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0, 0, fmt.Errorf("invalid value")
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, nil
}

func peak(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func eggbox(x, y float64) float64 {
	return math.Pow(2, math.Sin(x)) * math.Pow(2, math.Sin(y)) / 12
}

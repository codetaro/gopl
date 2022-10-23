// 1 foot = 0.3048 meters
package main

import (
	"fmt"
	"os"
	"strconv"
)

type Foot float64
type Meter float64

func (f Foot) String() string  { return fmt.Sprintf("%g Feet", f) }
func (m Meter) String() string { return fmt.Sprintf("%g Meters", m) }

func FToM(f Foot) Meter { return Meter(f * 0.3048) }
func MToF(m Meter) Foot { return Foot(m / 0.3048) }

func main() {
	val, err := strconv.ParseFloat(os.Args[1], 64)
	unit := os.Args[2]
	if err != nil {
		fmt.Fprintf(os.Stderr, "fm: %v\n", err)
		os.Exit(1)
	}
	switch unit {
	case "foot":
		f := Foot(val)
		m := FToM(f)
		fmt.Printf("%s = %s\n", f, m)
	case "meter":
		m := Meter(val)
		f := MToF(m)
		fmt.Printf("%s = %s\n", m, f)
	default:
		fmt.Printf("Unsupport unit: %v\n", unit)
	}
}

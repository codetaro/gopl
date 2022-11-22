package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

// comma inserts commas in a non-negative decimal integer string
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func comma2(s string) string {
	// sign
	var sign string
	if unicode.IsDigit(rune(s[0])) {
		sign = ""
	} else {
		sign = string(s[0])
		s = s[1:]
	}
	//
	dot := strings.LastIndex(s, ".")
	var integral, fractional string
	if dot >= 0 {
		integral = s[:dot]
		fractional = s[dot:]
	} else {
		integral = s
		fractional = ""
	}

	var buf bytes.Buffer
	for i := 0; i < len(integral); i++ {
		buf.WriteByte(integral[i])
		if (len(integral)-1-i)%3 == 0 &&
			i != len(s)-1 &&
			i != dot-1 {
			buf.WriteString(",")
		}
	}
	return sign + buf.String() + fractional
}

func main() {
	s1 := "1234567"
	s2 := "+1234.567"
	s3 := "-123.4567"
	fmt.Println(comma2(s1))
	fmt.Println(comma2(s2))
	fmt.Println(comma2(s3))
}

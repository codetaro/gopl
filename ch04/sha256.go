package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func main() {
	var in []byte
	var flag string
	fmt.Scanln(&in)
	fmt.Scanln(&flag)

	sha(in, flag)
}

// Exercise 4.1
func diff(c1, c2 [32]byte) int {
	var count int
	for i := 0; i < 32; i++ {
		if c1[i] != c2[i] {
			count++
		}
	}
	return count
}

// Exercise 4.2
func sha(in []byte, flag string) {
	switch flag {
	case "384":
		fmt.Println(sha512.Sum384(in))
	case "512":
		fmt.Println(sha512.Sum512(in))
	default:
		fmt.Println(sha256.Sum256(in))
	}
}

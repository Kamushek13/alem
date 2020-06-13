package main

import "github.com/01-edu/z01"

func main() {
	for start := byte('0'); start <= byte('9'); start++ {
		z01.PrintRune(rune(start))
	}
	z01.PrintRune('\n')
}

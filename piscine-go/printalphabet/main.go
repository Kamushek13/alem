package main

import "github.com/01-edu/z01"

func main() {
	for start := byte('a'); start <= byte('z'); start++ {
		z01.PrintRune(rune(start))
	}
	z01.PrintRune('\n')
}

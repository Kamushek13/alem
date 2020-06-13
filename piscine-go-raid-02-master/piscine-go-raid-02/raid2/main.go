package main

import (
	"fmt"
	"os"
)

func main() {
	template := os.Args[1:]
	if !templateValidity(template) {
		return
	}
	var newarr [9][9]int
	newarr = convertIntoArr(template)

	if !solvedsudoku(newarr) {
		fmt.Println("Error")
	}
}

func solvedsudoku(newarr [9][9]int) bool {

	var row, col int
	isFilled := true

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if newarr[i][j] == 0 {
				row = i
				col = j
				isFilled = false
				break
			}
		}
		if !isFilled {
			break
		}
	}

	if isFilled {
		for i := 0; i < 9; i++ {
			for index, j := range newarr[i] {
				fmt.Print(j)
				if index != 8 {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}

		return true
	}

	for n := 1; n <= 9; n++ {
		if possible(newarr, row, col, n) {
			newarr[row][col] = n
			if solvedsudoku(newarr) {
				return true
			}
			newarr[row][col] = 0
		}
	}
	return false
}

func possible(newarr [9][9]int, row int, col int, num int) bool {
	for i := 0; i < 9; i++ {
		if newarr[row][i] == num || newarr[i][col] == num {
			return false
		}
	}

	sqrt := 3
	rowstart := row - row%sqrt
	colstart := col - col%sqrt

	for i := rowstart; i < rowstart+sqrt; i++ {
		for j := colstart; j < colstart+sqrt; j++ {
			if newarr[i][j] == num {
				return false
			}
		}
	}
	return true
}

func convertIntoArr(newarr []string) [9][9]int {
	var arr [9][9]int
	var count int
	for i, line := range newarr {
		for j, value := range line {
			if value == '.' {
				arr[i][j] = 0
			} else {
				count = 0
				for nb := '0'; nb < value; nb++ {
					count++
				}
				arr[i][j] = count
			}
		}
	}
	return arr
}

func templateValidity(s []string) bool {
	if len(s) != 9 {
		fmt.Println("Error")
		return false
	}
	for _, line := range s {
		if len(line) != 9 {
			fmt.Println("Error")
			return false
		}
		for _, value := range line {
			if value != '.' && (value < '1' || value > '9') {
				fmt.Println("Error")
				return false
			}
		}
	}
	return true
}

package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	number := []int{}

	count := 0
	_, number, count = recursiveProver(n, number, count)
	number = bubbleSort(number, count)

	for _, n := range number {
		z01.PrintRune(rune(n + 48))

	}
}

func recursiveProver(m int, number []int, count int) (int, []int, int) {
	count++
	number = append(number, m%10)
	if m/10 == 0 { //n -??
		return m, number, count
	}
	return recursiveProver(m/10, number, count)
}

func bubbleSort(a []int, N int) []int { // ??

	for i := 0; i < N; i++ {
		for j := 0; j < N-1; j++ {
			if a[j] > a[j+1] {
				aa := a[j]
				a[j] = a[j+1]
				a[j+1] = aa
			}
		}
	}
	return a
}

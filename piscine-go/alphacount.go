package piscine

func AlphaCount(str string) int {
	nb := 0
	for _, letter := range str {
		if (letter >= 97 && letter <= 122) || (letter >= 65 && letter <= 90) {
			nb++
		}
	}
	return nb
}

package piscine

func TrimAtoi(s string) int {
	var str []rune
	number := 0
	degree := 1
	index := 0
	isFirst := true
	isNegative := false
	for i := range s {
		if rune(s[i]) == 45 {
			if isFirst {
				isNegative = true
				isFirst = false
			}
		}
		if rune(s[i]) >= 48 && rune(s[i]) <= 57 {
			isFirst = false
			str = append(str, rune(s[i]))
			index++
		}
	}
	for i := 1; i <= index; i++ {
		number += int(str[index-i]-48) * degree
		degree *= 10
	}
	if isNegative {
		number *= -1
	}
	return number
}

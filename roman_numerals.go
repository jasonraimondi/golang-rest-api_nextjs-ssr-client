package main

func RomanNumeral(s string) int {
	if s == "I" {
		return 1
	}
	result := 0
	for _, r := range s {
		if string(r) == "I" {
			result += 1
		}
	}
	return result
}

package kata

func RomanNumeral(s string) int {
	if s == "I" {
		return 1
	}
	result := 0
	strlen := len(string(s))
	for i, v := range s {
		currentLetter := string(v)
		hasNextLetter := strlen > i+1
		var nextLetter string
		if hasNextLetter {
			nextLetter = string([]rune(s)[i+1])
		}

		if currentLetter == "I" && hasNextLetter && nextLetter == "V" {
			result += 4
		} else if currentLetter == "I" && hasNextLetter && nextLetter == "X" {
			result += 9
		} else if currentLetter == "X" && hasNextLetter && nextLetter == "L" {
			result += 40
		} else if currentLetter == "X" && hasNextLetter && nextLetter == "C" {
			result += 90
		} else if currentLetter == "I" {
			result += 1
		} else if currentLetter == "V" {
			result += 5
		} else if currentLetter == "X" {
			result += 10
		} else if currentLetter == "L" {
			result += 50
		} else if currentLetter == "C" {
			result += 100
		}
	}
	return result
}

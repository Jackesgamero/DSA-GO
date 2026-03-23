package strings

func ReverseVowels(str string) (string, error) {
	runes := []rune(str)
	i, j := 0, len(runes)-1

	for i < j {
		if !isVowel(runes[i]) {
			i++
			continue
		}

		if !isVowel(runes[j]) {
			j--
			continue
		}

		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}
	return string(runes), nil
}

func isVowel(c rune) bool {
	switch c {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	}
	return false
}

package string_arrray

func ReverseVowels(s string) string {
	bytes := []byte(s)
	left, right := 0, len(bytes)-1

	for left < right {
		for left < right && !isVowel(bytes[left]) {
			left++
		}
		for left < right && !isVowel(bytes[right]) {
			right--
		}

		if left < right {
			bytes[left], bytes[right] = bytes[right], bytes[left]
			left++
			right--
		}
	}
	return string(bytes)
}

func isVowel(char byte) bool {
	switch char {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	case 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}

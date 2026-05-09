package main

// optimal solution
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var count [26]int
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}

	for _, v := range count {
		if v != 0 {
			return false
		}
	}

	return true
}

// my solution
/*
	func isAnagram(s string, t string) bool {
		if len(s) != len(t) {
			return false
		}

		myChars := make(map[byte]int)

		for i := 0; i < len(s); i++ {
			if v, ok := myChars[s[i]]; !ok {
				myChars[s[i]] = 1
			} else {
				v++
				myChars[s[i]] = v
			}
		}

		for i := 0; i < len(t); i++ {
			if v, ok := myChars[t[i]]; !ok {
				return false
			} else {
				v--
				myChars[t[i]] = v
			}
		}

		for _, v := range myChars {
			if v != 0 {
				return false
			}
		}
		return true
	}
*/

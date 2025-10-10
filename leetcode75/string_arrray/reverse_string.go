package string_arrray

import "strings"

func reverseWords(s string) string {
	s = strings.TrimSpace(s)

	parts := strings.Split(s, " ")

	for i := 0; i < len(parts); i++ {
		parts[i] = strings.TrimSpace(parts[i])
	}

	var b strings.Builder

	for i := len(parts) - 1; i >= 0; i-- {
		b.WriteString(parts[i])

		if i > 0 {
			b.WriteString(" ")
		}
	}

	return b.String()
}

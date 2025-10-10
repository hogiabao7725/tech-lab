package string_arrray

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	len_min := min(len(word1), len(word2))
	var x strings.Builder

	for i := 0; i < len_min; i++ {
		x.WriteByte(word1[i])
		x.WriteByte(word2[i])
	}

	if len_min == len(word1) {
		x.WriteString(word2[len_min:])
	} else {
		x.WriteString(word1[len_min:])
	}

	return x.String()
}

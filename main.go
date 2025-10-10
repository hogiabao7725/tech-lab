package main

import (
	"fmt"
	"strings"
)

func main() {
	parts := strings.Split("a good   example", " ")
	fmt.Println(parts)

	parts2 := strings.Fields("a good   example")
	fmt.Println(parts2)
}

package main

import (
	"bufio"
	"os"
	"strings"
)

var romanMap = initializeRomanMap()
var intToRomanMap = intToRoman()
var a, b *int
var operators = map[string]func() int{
	"+": func() int { return *a + *b },
	"-": func() int { return *a - *b },
	"/": func() int { return *a / *b },
	"*": func() int { return *a * *b },
}
var data []string

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		console, _ := reader.ReadString('\n')
		s := strings.ReplaceAll(console, " ", "")
		calculator(strings.ToUpper(strings.TrimSpace(s)))
	}
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func calculator(s string) {
	var operator string
	var stringsFound int
	numbers := make([]int, 0)
	romans := make([]string, 0)
	romansToInt := make([]int, 0)
	for idx := range operators {
		for _, val := range s {
			if idx == string(val) {
				operator += idx
				data = strings.Split(s, operator)
			}
		}
	}
	switch {
	case len(operator) > 1:
		panic("математический оператор не верный")
	case len(operator) < 1:
		panic("введенная строка не является математической операцией")
	}
	for _, elem := range data {
		num, err := strconv.Atoi(elem)
		if err != nil {
			stringsFound++
			romans = append(romans, elem)
		} else {
			numbers = append(numbers, num)
		}
	}

	switch stringsFound {
	case 1:
		panic("используются разные системы счисления")
	case 0:
		errCheck := numbers[0] > 0 && numbers[0] < 11 &&
			numbers[1] > 0 && numbers[1] < 11
		if val, ok := operators[operator]; ok && errCheck == true {
			a, b = &numbers[0], &numbers[1]
			fmt.Println(val())
		} else {
			panic("необходимо ввести только арабские или римские цифры от 1 до 10 включительно")
		}
	case 2:
		for _, elem := range romans {
			if val, ok := romanMap[elem]; ok && val > 0 && val < 11 {
				romansToInt = append(romansToInt, val)
			} else {
				panic("необходимо ввести только арабские или римские цифры от 1 до 10 включительно")
			}
		}
		if val, ok := operators[operator]; ok {
			a, b = &romansToInt[0], &romansToInt[1]
			convIntToRoman(val())
		}
	}
}

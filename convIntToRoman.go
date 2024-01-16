package main

import "fmt"

func convIntToRoman(romanResult int) {
	var romanNum string
	if romanResult == 0 {
		panic("в римской системе нет числа 0")
	} else if romanResult < 0 {
		panic("в римской системе нет отрицательных чисел")
	}
	for romanResult > 0 {
		for _, elem := range intToRomanMap {
			for i := elem; i <= romanResult; {
				for index, value := range romanMap {
					if value == elem {
						romanNum += index
						romanResult -= elem
					}
				}
			}
		}
	}
	fmt.Println(romanNum)
}

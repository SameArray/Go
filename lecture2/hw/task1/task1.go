package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Print("Введите число: ")
	// fmt.Scan(&m)
	var m int = 1417
	result := intToRoman(m)
	fmt.Print(result)
}

func intToRoman(num int) string {
	m := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	var result string

	for _, i := range nums {
		if num >= i {
			r := num / i
			result = result + strings.Repeat(m[i], r)
			num = num % i
		}
	}
	return result
}

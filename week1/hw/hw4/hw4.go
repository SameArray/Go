package main

import (
	"fmt"
	"sort"
)

// IntegerSlice представляет структуру для сортировки слайса целых чисел.
type IntegerSlice []int

// Len возвращает длину слайса.
func (s IntegerSlice) Len() int {
	return len(s)
}

// Less определяет, должен ли элемент с индексом i считаться меньшим, чем элемент с индексом j.
func (s IntegerSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

// Swap меняет местами элементы с индексами i и j.
func (s IntegerSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	// Создаем слайс целых чисел
	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}

	// Используем нашу структуру для сортировки
	sort.Sort(IntegerSlice(numbers))

	// Выводим отсортированный слайс
	fmt.Println(numbers)
}

package main

import (
	"testing"
)

func TestCompareSlice(t *testing.T) {
	slice1 := []int{9, 5, 4, 1, 7}
	slice2 := []int{4, 5, 1, 7, 9}
	result := CompareSlice(slice1, slice2)
	if !result {
		t.Errorf("Ожидалось, что слайсы совпадают, но получено false")
	}

	slice3 := []int{1, 2, 3}
	slice4 := []int{1, 2, 1}
	result = CompareSlice(slice3, slice4)
	if result {
		t.Errorf("Ожидалось, что слайсы не совпадают, но получено true")
	}

	slice5 := []int{1, 2, 3}
	slice6 := []int{1, 2, 3, 4}
	result = CompareSlice(slice5, slice6)
	if result {
		t.Errorf("Ожидалось, что слайсы не совпадают из-за разной длины, но получено true")
	}

	slice7 := []int{}
	slice8 := []int{}
	result = CompareSlice(slice7, slice8)
	if !result {
		t.Errorf("Ожидалось, что пустые слайсы считаются совпадающими, но получено false")
	}
}

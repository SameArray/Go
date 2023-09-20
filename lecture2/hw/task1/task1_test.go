package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	var m1 int = 1
	result1 := intToRoman(m1)
	if result1 != "I" {
		t.Errorf("Тестовый случай 1 не пройден: ожидалось 'I', но получено '%s'", result1)
	}

	var m2 int = -4
	result2 := intToRoman(m2)
	if result2 != "" {
		t.Errorf("Тестовый случай 2 не пройден: ожидалось 'nil', но получено '%s'", result2)
	}

	var m3 int = 0
	result3 := intToRoman(m3)
	if result3 != "" {
		t.Errorf("Тестовый случай 3 не пройден: ожидалось '0', но получено '%s'", result3)
	}

}

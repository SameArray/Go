package main

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	strs1 := []string{"flower", "flow", "flight"}
	result1 := longestCommonPrefix(strs1)
	if result1 != "fl" {
		t.Errorf("Тестовый случай 1 не пройден: ожидалось 'fl', но получено '%s'", result1)
	}

	strs2 := []string{"dog", "racecar", "car"}
	result2 := longestCommonPrefix(strs2)
	if result2 != "" {
		t.Errorf("Тестовый случай 2 не пройден: ожидалось '', но получено '%s'", result2)
	}

	strs3 := []string{}
	result3 := longestCommonPrefix(strs3)
	if result3 != "" {
		t.Errorf("Тестовый случай 3 не пройден: ожидалось '', но получено '%s'", result3)
	}

	strs4 := []string{"apple"}
	result4 := longestCommonPrefix(strs4)
	if result4 != "apple" {
		t.Errorf("Тестовый случай 4 не пройден: ожидалось 'apple', но получено '%s'", result4)
	}

	strs5 := []string{"abcdefg", "abcde", "abc"}
	result5 := longestCommonPrefix(strs5)
	if result5 != "abcde" {
		t.Errorf("Тестовый случай 5 не пройден: ожидалось 'abcde', но получено '%s'", result5)
	}
}

package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		j := 0
		for j < len(prefix) && j < len(strs[i]) && prefix[j] == strs[i][j] {
			j++
		}

		prefix = prefix[:j]

		if prefix == "" {
			return ""
		}
	}

	return prefix
}

func main() {
	strs1 := []string{"flower", "flow", "flight"}
	result1 := longestCommonPrefix(strs1)
	fmt.Println("Output 1:", result1)

	strs2 := []string{"dog", "racecar", "car"}
	result2 := longestCommonPrefix(strs2)
	fmt.Println("Output 2:", result2)
}

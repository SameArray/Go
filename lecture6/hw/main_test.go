package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGenerateStructFromJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "simple",
			input:    `{"name": "Nurzhan", "age": 20}`,
			expected: "type MyStruct struct {\nAge float64 `json:\"age\"`\nName string `json:\"name\"`\n}\n",
		},
		{
			name:     "nested",
			input:    `{"user": {"name": "Nurzhan", "age": 20}}`,
			expected: "type MyStructUser struct {\nAge float64 `json:\"age\"`\nName string `json:\"name\"`\n}\ntype MyStruct struct {\nUser MyStructUser `json:\"user\"`\n}\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateStructFromJSON("MyStruct", tt.input)
			if err != nil {
				t.Fatalf("error: %v", err)
			}

			if diff := cmp.Diff(tt.expected, got); diff != "" {
				t.Fatalf("mismatch (-expected +got):\n%s", diff)
			}
		})
	}
}

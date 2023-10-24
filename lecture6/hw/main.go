package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func GenerateStructFromJSON(structName, jsonStr string) (string, error) {
	var data interface{}
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		return "", err
	}

	var fields []string

	processData := func(data map[string]interface{}, structName string, depth int) string {
		if depth > 5 {
			return ""
		}

		fields = []string{}

		for key, value := range data {
			fieldType := getType(value)
			if fieldType == "object" {
				nestedStructName := structName + strings.Title(key)
				fieldType = nestedStructName
				fields = append(fields, processData(value.(map[string]interface{}), nestedStructName, depth+1))
			}
			fields = append(fields, fmt.Sprintf("%s %s `json:\"%s\"`", strings.Title(key), fieldType, key))
		}

		return fmt.Sprintf("type %s struct {\n%s\n}\n", structName, strings.Join(fields, "\n"))
	}

	return processData(data.(map[string]interface{}), structName, 1), nil
}

func generate(name string, data map[string]interface{}, level int) (string, error) {
	if level > 5 {
		return "", fmt.Errorf("max nesting level exceeded")
	}

	var fields []string
	for key, val := range data {
		var fieldType string
		switch v := val.(type) {
		case map[string]interface{}:
			nestedStructName := fmt.Sprintf("%s%s", name, strings.Title(key))
			nestedStruct, err := generate(nestedStructName, v, level+1)
			if err != nil {
				return "", err
			}
			fields = append(fields, nestedStruct)
			fieldType = nestedStructName
		case float64:
			fieldType = "float64"
		case string:
			fieldType = "string"
		case bool:
			fieldType = "bool"
		default:
			return "", fmt.Errorf("unsupported type: %T", v)
		}
		fields = append(fields, fmt.Sprintf("%s %s `json:\"%s\"`", strings.Title(key), fieldType, key))
	}
	return fmt.Sprintf("type %s struct {\n%s\n}\n", name, strings.Join(fields, "\n")), nil
}

func main() {
	result := GenerateStructFromJSON("MyStruct", test.input)
	if !strings.EqualFold(test.expected, result) {
		t.Errorf("expected:\n%s\ngot:\n%s", test.expected, result)
	}
}

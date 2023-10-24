package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func GenerateStructFromJSON(name string, data string) (string, error) {
	var parsedData map[string]interface{}
	err := json.Unmarshal([]byte(data), &parsedData)
	if err != nil {
		return "", err
	}

	return generate(name, parsedData, 1)
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
	jsonStr := `{"name": "Nurzhan", "age": 20, "address": {"city": "Almaty"}}`
	generatedCode, err := GenerateStructFromJSON("Person", jsonStr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(generatedCode)
}

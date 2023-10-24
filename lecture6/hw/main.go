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
	return fmt.Sprintf("type %s struct {\n\t%s\n}\n", name, strings.Join(fields, "\n\t")), nil
}

func main() {
	inputJSON := `{"name": "Nurzhan", "age": 20}`
	expectedStruct := `type MyStruct struct {
	Name string ` + "`json:\"name\"`" + `
	Age float64 ` + "`json:\"age\"`" + `
}`

	result, err := GenerateStructFromJSON("MyStruct", inputJSON)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if !strings.EqualFold(expectedStruct, result) {
		fmt.Printf("expected:\n%s\ngot:\n%s", expectedStruct, result)
	} else {
		fmt.Println("Struct generated successfully:")
		fmt.Println(result)
	}
}

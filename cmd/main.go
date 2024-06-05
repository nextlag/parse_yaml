package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func parseYAML(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("ошибка чтения файла: %v", err)
	}

	var parsedData map[string]interface{}
	err = yaml.Unmarshal(data, &parsedData)
	if err != nil {
		yamlError, ok := err.(*yaml.TypeError)
		if ok && len(yamlError.Errors) > 0 {
			return fmt.Errorf("ошибка парсинга YAML: %v", yamlError.Errors[0])
		}
		return fmt.Errorf("ошибка парсинга YAML: %v", err)
	}

	fmt.Println("YAML успешно распарсен:", parsedData)
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("пожалуйста, укажите путь к YAML файлу")
		return
	}
	filePath := os.Args[1]
	err := parseYAML(filePath)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
}

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func parseYAML(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, fmt.Errorf("ошибка открытия файла: %v", err)
	}
	defer file.Close()

	var lineCount int
	reader := bufio.NewReader(file)
	var data []byte

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return 0, fmt.Errorf("ошибка чтения файла: %v", err)
		}
		lineCount++
		data = append(data, line...)
	}

	var parsedData map[string]interface{}
	err = yaml.Unmarshal(data, &parsedData)
	if err != nil {
		// Попробуем найти строку, на которой произошла ошибка
		var yamlErr *yaml.TypeError
		if ok := errors.As(err, &yamlErr); ok {
			return lineCount, fmt.Errorf("ошибка парсинга YAML: %v", yamlErr.Errors)
		}
		return lineCount, fmt.Errorf("ошибка парсинга YAML: %v", err)
	}

	return lineCount, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("пожалуйста, укажите путь к YAML файлу")
		return
	}
	filePath := os.Args[1]
	lineCount, err := parseYAML(filePath)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Printf("Файл успешно распарсен, количество строк: %d\n", lineCount)
}

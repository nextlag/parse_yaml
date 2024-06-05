package main

import (
	"bufio"
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
	scanner := bufio.NewScanner(file)
	var data []byte

	for scanner.Scan() {
		lineCount++
		data = append(data, scanner.Bytes()...)
		data = append(data, '\n')
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("ошибка чтения файла: %v", err)
	}

	var parsedData map[string]interface{}
	decoder := yaml.NewDecoder(bufio.NewReader(file))
	if err := decoder.Decode(&parsedData); err != nil {
		return lineCount, fmt.Errorf("ошибка парсинга YAML на строке %d: %v", lineCount, err)
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

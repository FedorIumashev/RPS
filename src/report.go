package src

import (
	"fmt"
	"os"
	"strings"
)

// Структура Issue для хранения информации о найденных проблемах
type Issue struct {
	FilePath string
	Severity string
	Message  string
	Line     int
}

// Функция сканирования с учётом типа файла
func RunScan(path string) []Issue {
	var issues []Issue

	// Чтение содержимого файла
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return issues
	}

	// Примерная логика сканирования
	var severity, message string
	lines := strings.Split(string(content), "\n")

	// Проверяем тип файла и находим проблему в зависимости от содержимого
	switch {
	case strings.HasSuffix(path, ".py"):
		severity = "high"
		message = "Sample issue found in Python file"
		// Примерная проверка для Python файлов
		for i, line := range lines {
			if strings.Contains(line, "print") {
				issues = append(issues, Issue{
					FilePath: path,
					Severity: severity,
					Message:  "Potential issue: usage of print() found in Python",
					Line:     i + 1, // Номерация строк начинается с 1
				})
			}
		}

	case strings.HasSuffix(path, ".go"):
		severity = "medium"
		message = "Sample issue found in Go file"
		// Примерная проверка для Go файлов
		for i, line := range lines {
			if strings.Contains(line, "fmt.Println") {
				issues = append(issues, Issue{
					FilePath: path,
					Severity: severity,
					Message:  "Potential issue: fmt.Println() found in Go",
					Line:     i + 1, // Номерация строк начинается с 1
				})
			}
		}

	default:
		severity = "low"
		message = "Sample issue found in unknown file type"
		issues = append(issues, Issue{
			FilePath: path,
			Severity: severity,
			Message:  message,
			Line:     10, // Пример строки, на которой обнаружена ошибка
		})
	}

	fmt.Println("Сканирование завершено для:", path)
	return issues
}

// Функция для генерации отчета
func GenerateReport(issues []Issue, outputPath string) {
	if len(issues) == 0 {
		fmt.Println("Ошибок не найдено!")
		return
	}

	// Создание файла отчета
	file, err := os.Create(outputPath)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	// Запись информации в файл отчета
	for _, issue := range issues {
		reportLine := fmt.Sprintf("Файл: %s\nСерьезность: %s\nСообщение: %s\nСтрока: %d\n\n", issue.FilePath, issue.Severity, issue.Message, issue.Line)
		_, err := file.WriteString(reportLine)
		if err != nil {
			fmt.Println("Ошибка при записи в файл отчета:", err)
			return
		}
	}

	fmt.Println("Отчет успешно сгенерирован:", outputPath)
}

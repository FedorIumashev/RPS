package src

import (
    "fmt"
    "os"
	"strings"
)

type Issue struct {
    FilePath string
    Severity string
    Message  string
    Line     int
}

// Функция сканирования с учётом типа файла и правильного номера строки
func RunScan(path string) []Issue {
    var issues []Issue

    // Чтение содержимого файла
    content, err := os.ReadFile(path)
    if err != nil {
        fmt.Println("Ошибка при чтении файла:", err)
        return issues
    }

    lines := strings.Split(string(content), "\n")

    var severity, message string

    // Проверяем тип файла и находим проблему в зависимости от содержимого
    switch {
    case path[len(path)-3:] == ".py":
        severity = "high"
        message = "Sample issue found in Python file"
        // Примерная проверка для Python файлов
        for i, line := range lines {
            if strings.Contains(line, "print") {
                message = "Potential issue: usage of print() found in Python"
                issues = append(issues, Issue{
                    FilePath: path,
                    Severity: severity,
                    Message:  message,
                    Line:     i + 1, // Номер строки, где найдено совпадение
                })
            }
        }

    case path[len(path)-3:] == ".go":
        severity = "medium"
        message = "Sample issue found in Go file"
        // Примерная проверка для Go файлов
        for i, line := range lines {
            if strings.Contains(line, "fmt.Println") {
                message = "Potential issue: fmt.Println() found in Go"
                issues = append(issues, Issue{
                    FilePath: path,
                    Severity: severity,
                    Message:  message,
                    Line:     i + 1, // Номер строки, где найдено совпадение
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
            Line:     10, 
        })
    }

    fmt.Println("Сканирование завершено для:", path)
    return issues
}


// Функция для генерации TXT отчета
func GenerateReportTxt(issues []Issue, outputPath string) {
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

// Функция для генерации CSV отчета
func GenerateReportCsv(issues []Issue, outputPath string) {
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

    // Запись заголовка CSV файла
    _, err = file.WriteString("File Path, Severity, Message, Line\n")
    if err != nil {
        fmt.Println("Ошибка при записи в файл отчета:", err)
        return
    }

    // Запись каждой ошибки в CSV файл
    for _, issue := range issues {
        csvLine := fmt.Sprintf("\"%s\", \"%s\", \"%s\", %d\n", issue.FilePath, issue.Severity, issue.Message, issue.Line)
        _, err := file.WriteString(csvLine)
        if err != nil {
            fmt.Println("Ошибка при записи в файл отчета:", err)
            return
        }
    }

    fmt.Println("Отчет успешно сгенерирован:", outputPath)
}

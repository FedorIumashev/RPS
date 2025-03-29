package src

import (
    "encoding/csv"
    "fmt"
    "os"
)

// Структура для хранения информации об ошибке
type Issue struct {
    FilePath string
    Severity string
    Message  string
    Line     int
}

// GenerateReport создает отчет в CSV-формате
func GenerateReport(issues []Issue) {
    if len(issues) == 0 {
        fmt.Println("Нет данных для отчета.")
        return
    }

    file, err := os.Create("report.csv")
    if err != nil {
        fmt.Printf("Ошибка создания файла: %v\n", err)
        return
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    // Заголовки
    err = writer.Write([]string{"File", "Severity", "Message", "Line"})
    if err != nil {
        fmt.Printf("Ошибка записи заголовков в CSV: %v\n", err)
        return
    }

    // Записываем каждую проблему
    for _, issue := range issues {
        err = writer.Write([]string{
            issue.FilePath,
            issue.Severity,
            issue.Message,
            fmt.Sprintf("%d", issue.Line),
        })
        if err != nil {
            fmt.Printf("Ошибка записи строки в CSV: %v\n", err)
            return
        }
    }

    fmt.Println("Отчёт сохранён как report.csv")
}

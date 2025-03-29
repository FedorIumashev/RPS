// src/report.go

package src

import "fmt"

type Issue struct {
    FilePath string
    Severity string
    Message  string
    Line     int
}

// Генерация отчета на основе найденных проблем
func GenerateReport(issues []Issue) {
    if len(issues) == 0 {
        fmt.Println("Ошибок не найдено!")
        return
    }

    fmt.Println("Обнаружены проблемы:")
    for _, issue := range issues {
        fmt.Printf("Файл: %s\nСерьезность: %s\nСообщение: %s\nСтрока: %d\n\n", 
                   issue.FilePath, issue.Severity, issue.Message, issue.Line)
    }

    // Если нужно, записывать отчет в файл или сохранять как CSV, TXT и т. д.
}

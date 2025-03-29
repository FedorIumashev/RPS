// src/rps.go

package src

import "fmt"

// Эта функция будет заниматься сканированием
func RunScan(path string) []Issue {
    var issues []Issue

    // Примерная логика сканирования
    issues = append(issues, Issue{
        FilePath: path,
        Severity: "high",
        Message:  "Sample issue found",
        Line:     10,
    })

    fmt.Println("Сканирование завершено для:", path)
    return issues
}

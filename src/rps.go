// Package src содержит функции для работы с репозиториями и сканирования кода
// на наличие уязвимостей. Включает в себя функции для генерации отчетов и  сканирования
// кода на наличие уязвимостей.
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

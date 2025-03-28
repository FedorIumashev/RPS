package src

import (
    "encoding/csv"
    "fmt"
    "os"
    //"path/filepath"
    //"strings"
)

func GenerateReport(data []string) {
    file, err := os.Create("report.csv")
    if err != nil {
        fmt.Println("Ошибка создания файла:", err)
        return
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    err = writer.Write(data)
    if err != nil {
        fmt.Println("Ошибка записи в CSV:", err)
        return
    }
    writer.Flush()
    fmt.Println("Отчёт сохранён как report.csv")
}
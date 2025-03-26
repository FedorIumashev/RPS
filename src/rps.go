package src

import (
    "fmt"
    "log"
    "path/filepath"
    "github.com/FedorIumashev/RPS/src/scan"  // Пакет для сканирования
)

func RunScan(directory string) {
    // Получаем абсолютный путь к каталогу
    absPath, err := filepath.Abs(directory)
    if err != nil {
        log.Fatalf("Ошибка получения абсолютного пути: %v\n", err)
    }
    fmt.Println("Анализируем каталог:", absPath)

    // Запуск Bandit
    err = scan.RunBandit(absPath)
    if err != nil {
        log.Printf("Ошибка при запуске Bandit: %v\n", err)
    } else {
        fmt.Println("Bandit завершил анализ")
    }

    // Запуск Nmap
    err = scan.RunNmap(absPath)
    if err != nil {
        log.Printf("Ошибка при запуске Nmap: %v\n", err)
    } else {
        fmt.Println("Nmap завершил сканирование")
    }
}

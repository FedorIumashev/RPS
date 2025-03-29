package main

import (
    "fmt"
    "log"
    "os"
    "github.com/spf13/cobra"
    "github.com/FedorIumashev/RPS/src"  // Импортируем пакет src
)

func main() {
    var rootCmd = &cobra.Command{Use: "rps"}

    // Команда help
    var helpCmd = &cobra.Command{
        Use:   "help",
        Short: "Показывает список доступных команд",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("Доступные команды:")
            fmt.Println("rps help - Показывает список доступных команд")
            fmt.Println("rps test <file> - Запускает тесты для указанного файла")
            fmt.Println("rps test <directory> - Запускает тесты для всех файлов в директории")
            fmt.Println("rps test-report <directory> - Генерирует отчет по результатам тестов")
        },
    }
    rootCmd.AddCommand(helpCmd)

    // Команда test для файла
    var testFileCmd = &cobra.Command{
        Use:   "test",
        Short: "Запускает тесты безопасности для файла",
        Run: func(cmd *cobra.Command, args []string) {
            if len(args) < 1 {
                log.Fatal("Не указан файл для тестирования")
            }
            file := args[0]
            fmt.Println("Запуск тестов для файла:", file)
            issues := src.RunScan(file) // Получаем ошибки для файла
            src.GenerateReport(issues) // Генерируем отчет
        },
    }
    rootCmd.AddCommand(testFileCmd)

    // Команда test для директории
    var testDirectoryCmd = &cobra.Command{
        Use:   "test",
        Short: "Запускает тесты безопасности для директории",
        Run: func(cmd *cobra.Command, args []string) {
            if len(args) < 1 {
                log.Fatal("Не указана директория для тестирования")
            }
            directory := args[0]
            fmt.Println("Запуск тестов для директории:", directory)
            issues := src.RunScan(directory) // Получаем ошибки для директории
            src.GenerateReport(issues) // Генерируем отчет
        },
    }
    rootCmd.AddCommand(testDirectoryCmd)

    // Команда test-report для создания отчета
    var testReportCmd = &cobra.Command{
        Use:   "test-report",
        Short: "Генерирует отчет по результатам тестов в TXT или CSV",
        Run: func(cmd *cobra.Command, args []string) {
            if len(args) < 1 {
                log.Fatal("Не указана директория для отчетов")
            }
            directory := args[0]
            fmt.Println("Генерация отчета для директории:", directory)
            issues := src.RunScan(directory) // Получаем ошибки для директории
            src.GenerateReport(issues) // Генерация отчета
        },
    }
    rootCmd.AddCommand(testReportCmd)

    // Запуск основной команды
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/FedorIumashev/RPS/src" // Импортируем пакет src
	"github.com/spf13/cobra"
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
			fmt.Println("rps test-file <file> - Запускает тесты для указанного файла")
			fmt.Println("rps test-dir <directory> - Запускает тесты для всех файлов в директории")
			fmt.Println("rps test-report <directory> --format <txt|csv> - Генерирует отчет по результатам тестов")
		},
	}
	rootCmd.AddCommand(helpCmd)

	var testFileCmd = &cobra.Command{
		Use:   "test-file",
		Short: "Запускает тесты безопасности для файла",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				log.Fatal("Не указан файл для тестирования")
			}
			file := args[0]
			fmt.Println("Запуск тестов для файла:", file)
			issues := src.RunScan(file)                   // Получаем ошибки для файла
			outputFilePath := "report.txt"                // Пример файла отчета
			src.GenerateReportTxt(issues, outputFilePath) // Генерация отчета в TXT
		},
	}
	rootCmd.AddCommand(testFileCmd)

	// Команда test для директории
	var testDirectoryCmd = &cobra.Command{
		Use:   "test-dir",
		Short: "Запускает тесты безопасности для директории",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				log.Fatal("Не указана директория для тестирования")
			}
			directory := args[0]
			fmt.Println("Запуск тестов для директории:", directory)

			var allIssues []src.Issue
			err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() { // Пропускаем директории
					issues := src.RunScan(path) // Получаем ошибки для каждого файла
					allIssues = append(allIssues, issues...)
				}
				return nil
			})

			if err != nil {
				log.Fatal("Ошибка при сканировании директории:", err)
			}

			outputFilePath := "report.txt"                   // Пример файла отчета
			src.GenerateReportTxt(allIssues, outputFilePath) // Генерация отчета в TXT
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

			// Получаем формат отчета из флага
			format, _ := cmd.Flags().GetString("format")
			if format == "" {
				format = "txt" // По умолчанию генерируем TXT
			}

			var allIssues []src.Issue
			err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() { // Пропускаем директории
					issues := src.RunScan(path) // Получаем ошибки для каждого файла
					allIssues = append(allIssues, issues...)
				}
				return nil
			})

			if err != nil {
				log.Fatal("Ошибка при сканировании директории:", err)
			}

			// Генерация отчета в зависимости от формата
			if format == "csv" {
				outputFilePath := "report.csv"
				src.GenerateReportCsv(allIssues, outputFilePath) // Генерация отчета в CSV
			} else {
				outputFilePath := "report.txt"
				src.GenerateReportTxt(allIssues, outputFilePath) // Генерация отчета в TXT
			}
		},
	}

	// Добавляем флаг --format
	testReportCmd.Flags().String("format", "txt", "Формат отчета (txt или csv)")

	rootCmd.AddCommand(testReportCmd)

	// Запуск основной команды
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

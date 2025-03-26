package main

import (
    "fmt"
    "log"
    "net/http"
    "flag"
    "github.com/FedorIumashev/RPS/src"  // Импортируем пакет src
)

func StartServer() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello, DevSecOps!")
    })
    fmt.Println("Сервер работает на http://localhost:5000")
    log.Fatal(http.ListenAndServe(":5000", nil))
}

func main() {
    // Флаг для выбора директории
    directory := flag.String("directory", ".", "Path to the directory to scan")
    flag.Parse()

    // Запуск сканирования
    src.RunScan(*directory)

    // Запуск веб-сервера
    go StartServer()

    fmt.Println("Программа завершена.")
}

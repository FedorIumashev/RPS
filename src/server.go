package src

import (
    "fmt"
    "log"
    "net/http"
)

func StartServer() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello, DevSecOps!")
    })
    fmt.Println("Сервер работает на http://localhost:5000")
    log.Fatal(http.ListenAndServe(":5000", nil))
}

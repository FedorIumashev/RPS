package src

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, DevSecOps!")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server is running on http://localhost:5000")
    log.Fatal(http.ListenAndServe(":5000", nil))
}

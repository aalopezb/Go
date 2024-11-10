package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World with Go!")
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        log.Fatal("PORT not set")
    }

    http.HandleFunc("/", handler)

    log.Printf("Starting server on port %s...\n", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}

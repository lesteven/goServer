package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World!")
}

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/", handler)

    fileServer := http.FileServer(http.Dir("./ui/"))

    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    log.Fatal(http.ListenAndServe(":8080", mux))
}

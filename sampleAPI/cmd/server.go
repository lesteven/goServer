package main

import (
    "log"
    "net/http"
)


func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/", Home)
    fileServer := http.FileServer(http.Dir("../ui/static/"))
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    log.Fatal(http.ListenAndServe(":8080", mux))
}

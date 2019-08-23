package main

import (
    "log"
    "net/http"
    "flag"
)


func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/", Home)
    fileServer := http.FileServer(http.Dir("../ui/static"))
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    // default is :3000
    // go run *.go -addr=":num" for diff port
    addr := flag.String("addr", ":3000", "HTTP port")
    flag.Parse()

    log.Fatal(http.ListenAndServe(*addr, mux))
}

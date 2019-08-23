package main

import (
    "log"
    "net/http"
    "flag"
)


func main() {
    mux := http.NewServeMux()

    port := flag.String("port", ":8080", "PORT")
    origin := flag.String("origin", "http://localhost:3000", "CORS")
    flag.Parse()

    mux.HandleFunc("/", HomeWithFlag(origin))
    fileServer := http.FileServer(http.Dir("../ui/static/"))
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    log.Fatal(http.ListenAndServe(*port, mux))
}

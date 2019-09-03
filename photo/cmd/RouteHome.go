package main

import (
    "net/http"
    "log"
    "html/template"
)

func Home(w http.ResponseWriter, r *http.Request) {

    ts, _ := template.ParseFiles("../ui/html/index.html")
    err := ts.Execute(w, nil)

    if err != nil {
        log.Fatal("front server home ", err)
    }
}


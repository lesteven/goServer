package main

import (
    "net/http"
    "github.com/lesteven/goRes"
)

func Data(w http.ResponseWriter, r *http.Request) {
    goRes.SendJson(w, 200, "hello world")
}

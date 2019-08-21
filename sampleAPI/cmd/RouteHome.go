package main

import (
    "net/http"
    "github.com/lesteven/goRes"
)

type Person struct {
    FirstName string `json:"firstName"`
    LastName string `json:"lastName"`
    Occupation string `json:"occupation"`
    City string `json:"city"`
    Email string `json:"email"`
    Image string `json:"image"`
}

func Home(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin","*")
    goRes.SendJson(w, 200, data)
}

var url = "localhost:8080/static/"
var data = []Person{
    Person{"Steven", "Le", "Developer", "Oakland",
        "steven@example.com", url + "man1.svg"},
    Person{"Joe", "Schmoe", "Writer", "San Francisco",
        "joe@example.com", url + "man2.svg"},
    Person{"Plain", "Jane", "Developer", "San Jose",
        "jane@example.com", url + "woman1.svg"},
    Person{"Jon", "Smith", "Recruiter", "Davis",
        "jon@example.com", url + "man3.svg"},
    Person{"Kate", "Fate", "Analyst", "Milpitas",
        "kate@example.com", url + "woman2.svg"},
    Person{"Wayne", "Layne", "Lawyer", "Lawville",
        "wayne@example.com", url + "man4.svg"},
    Person{"Abbey", "Labby", "CEO", "Berkeley",
        "abbey@example.com", url + "woman3.svg"},
    Person{"Max", "Payne", "Hitman", "Painville",
        "max@example.com", url + "man5.svg"},
    Person{"Blake", "Snake", "Actor", "Los Angeles",
        "blake@example.com", url + "woman4.svg"},
    Person{"Neo", "One", "Scientist", "Matrix",
        "neo@example.com", url + "man6.svg"},
    Person{"Nicole", "Jones", "Docter", "Stockton",
        "nicole@example.com", url + "woman5.svg"},
    Person{"Golem", "Stone", "Comedian", "Modesto",
        "golem@example.com", url + "woman6.svg"},
}

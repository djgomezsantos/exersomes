// src/api.go
package main

import (
    "net/http"
    "encoding/json"
    "github.com/djgomez8/exersomes/molecular_types"
)

func ExerkineHandler(w http.ResponseWriter, r *http.Request) {
    exerkines := GetAllExerkines() // Would load from database/files
    json.NewEncoder(w).Encode(exerkines)
}

func NetworkHandler(w http.ResponseWriter, r *http.Request) {
    tissue := r.URL.Query().Get("tissue")
    exerkines := GetExerkinesByTissue(tissue)
    network := BuildNetwork(exerkines)
    json.NewEncoder(w).Encode(network)
}

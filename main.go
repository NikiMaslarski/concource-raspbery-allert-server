package main

import (
    "io"
    "net/http"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", handleBrokenTests)
    http.ListenAndServe(":8080", mux)
}

func handleBrokenTests(w http.ResponseWriter, r *http.Request){
    io.WriteString(w, "Hello World")
}

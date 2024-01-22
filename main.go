package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Fatal("Error parsing form: ", err)
	}

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Address: %s\n", address)

	return
}

func cringeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/cringe" {
		http.Error(w, "Not found", http.StatusNotFound)
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	fmt.Fprintf(w, "Cringe")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)

	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/cringe", cringeHandler)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("Error running server: ", err)
	}
}

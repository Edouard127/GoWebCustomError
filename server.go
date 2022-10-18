package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	// Create a web server
	server := mux.NewRouter()

	server.HandleFunc("/{status:.+}", StatusHandler).Methods("GET")
	server.HandleFunc("/", StatusHandler).Methods("GET")

	log.Printf("Listening on port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", server))
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	status := vars["status"]

	file := fmt.Sprintf("pages/%s.html", status)
	if _, err := os.Stat(file); err == nil {
		http.ServeFile(w, r, file)
	} else {
		http.ServeFile(w, r, "pages/404.html")
	}

	log.Printf("New request from %s, Method %s, Path %s", r.RemoteAddr, r.Method, r.URL)
}

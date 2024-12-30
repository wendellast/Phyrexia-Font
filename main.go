package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/index.html")
	})

	http.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/game.html")
	})

	portEnv := os.Getenv("PORT_PHYREXIA")
	port := "0.0.0.0"
	fullserver := port + ":" + portEnv

	log.Print("ServerInit")

	log.Fatal(http.ListenAndServe(fullserver, nil))

}

package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/map", mapHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		slog.Error("http listen and serve", "error", err)
		os.Exit(1)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mini App для выбора точки на карте")
}

func mapHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/map.html")
}

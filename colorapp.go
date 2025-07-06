package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func requestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s\n", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	color := os.Getenv("COLOR")
	if color == "" {
		color = "None"
	}

	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s\n", color)
	})

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("GET /coffee", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTeapot)
		fmt.Fprint(w, "The requested entity body is short and stout.")
	})

	addr := ":8080"

	log.Printf("Server listening on %v", addr)

	log.Printf("COLOR is set to %s", color)

	server := http.Server{
		Addr:    addr,
		Handler: requestLogger(mux),
	}

	log.Fatal(server.ListenAndServe())
}

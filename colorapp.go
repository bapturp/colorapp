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
	color := os.Getenv("COLOR")
	if color == "" {
		color = "None"
	}

	log.Printf("COLOR is set to %s", color)

	mux := http.NewServeMux()

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

	server := http.Server{
		Addr:    ":8080",
		Handler: requestLogger(mux),
	}

	log.Printf("Server listening on %v", server.Addr)

	log.Fatal(server.ListenAndServe())
}

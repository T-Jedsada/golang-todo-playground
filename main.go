package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/20scoops/todo-crud-go-playgound/index"
	"github.com/20scoops/todo-crud-go-playgound/todo"
	"github.com/gorilla/mux"
)

func main() {

	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	r := mux.NewRouter()
	r.Use(loggingMiddleware, mux.CORSMethodMiddleware(r))

	r.HandleFunc("/", index.IndexHandler)
	r.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	todoRoute := r.PathPrefix("/api/todo").Subrouter()
	todoRoute = todo.InitTodoRouters(todoRoute)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		ErrorLog:     logger,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	func() {
		log.Println("Starting Server -> localhost:8080")
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r)
		next.ServeHTTP(w, r)
	})
}

package todo

import "github.com/gorilla/mux"

func InitTodoRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/", createTodo).Methods("POST")
	router.HandleFunc("/", getAllTodos).Methods("GET")
	router.HandleFunc("/{todoID}", getATodo).Methods("GET")
	router.HandleFunc("/{todoID}", updateTodo).Methods("UPDATE")
	router.HandleFunc("/{todoID}", deleteTodo).Methods("DELETE")
	return router
}

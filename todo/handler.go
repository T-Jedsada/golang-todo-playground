package todo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/20scoops/todo-crud-go-playgound/models"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	todosResource struct {
		Data []models.Todo `json:"data"`
	}

	todoResource struct {
		Data models.Todo `json:"data"`
	}

	appError struct {
		Error      string `json:"error"`
		Message    string `json:"message"`
		HttpStatus int    `json:"status"`
	}

	errorResource struct {
		Data appError `json:"data"`
	}
)

func getAllTodos(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("todo")

	var todos []models.Todo
	iter := c.Find(nil).Iter()
	result := models.Todo{}
	for iter.Next(&result) {
		todos = append(todos, result)
	}
	j, err := json.Marshal(todosResource{Data: todos})
	if err != nil {
		displayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func getATodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["todoID"]

	context := NewContext()
	defer context.Close()
	c := context.DbCollection("todo")

	var todo models.Todo
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&todo)

	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			displayAppError(w, err, "An unexpected error has occurred", 500)
			return
		}
	}

	j, err := json.Marshal(todoResource{Data: todo})
	if err != nil {
		displayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var dataResource todoResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)

	if err != nil {
		displayAppError(w, err, "Invalid Todo data", 500)
		return
	}

	context := NewContext()
	defer context.Close()
	c := context.DbCollection("todo")

	todo := &dataResource.Data

	objID := bson.NewObjectId()
	todo.Id = objID
	todo.CreatedOn = time.Now()
	errInsert := c.Insert(&todo)

	if errInsert != nil {
		displayAppError(w, err, "Something went wrong", 500)
		return
	}

	j, err := json.Marshal(dataResource)
	if err != nil {
		displayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hellow World!")
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["todoID"]

	context := NewContext()
	defer context.Close()
	c := context.DbCollection("todo")

	err := c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})

	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			displayAppError(w, err, "An unexpected error has occurred", 500)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}

func displayAppError(w http.ResponseWriter, handlerError error, message string, code int) {
	errObj := appError{
		Error:      handlerError.Error(),
		Message:    message,
		HttpStatus: code,
	}
	log.Printf("AppError]: %s\n", handlerError)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(errorResource{Data: errObj}); err == nil {
		w.Write(j)
	}
}

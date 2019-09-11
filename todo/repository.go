package todo

import (
	"time"

	"github.com/20scoops/todo-crud-go-playgound/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TodoRepository struct {
	C *mgo.Collection
}

func (r *TodoRepository) getAllTodos() []models.Todo {
	var todos []models.Todo
	iter := r.C.Find(nil).Iter()
	result := models.Todo{}
	for iter.Next(&result) {
		todos = append(todos, result)
	}

	return todos
}

func (r *TodoRepository) getATodo(id string) (todo models.Todo, err error) {
	err = r.C.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&todo)
	return
}

func (r *TodoRepository) createTodo(todo *models.Todo) error {
	objID := bson.NewObjectId()
	todo.Id = objID
	todo.CreatedOn = time.Now()
	err := r.C.Insert(&todo)
	return err
}

func (r *TodoRepository) deleteTodo(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

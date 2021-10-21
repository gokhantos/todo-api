package repository

import (
	"context"
	"todo-api/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	GetTasks() interface{}
	FindTask(id string) interface{}
	CreateTask(i interface{}) interface{}
	UpdateTask(i interface{}) interface{}
	DeleteTask(id string)
}

type TodoRepository struct {
	collection *mongo.Collection
}

func NewTodoRepository(collection *mongo.Collection) Repository {
	return TodoRepository{collection: collection}
}

func (r TodoRepository) GetTasks() interface{} {
	var todos = make([]model.Todo, 0)
	found, _ := r.collection.Find(context.Background(), bson.D{})
	//fmt.Println(found)
	_ = found.All(context.Background(), &todos)
	//fmt.Println(todos)
	return todos
}

func (r TodoRepository) FindTask(id string) interface{} {
	var todo model.Todo
	filter := bson.M{"_id": id}
	one := r.collection.FindOne(context.Background(), filter)
	_ = one.Decode(&todo)
	return todo
}

func (r TodoRepository) CreateTask(i interface{}) interface{} {
	id, _ := r.collection.InsertOne(context.Background(), i)
	filter := bson.M{"_id": id.InsertedID}
	one := r.collection.FindOne(context.Background(), filter)
	var todo model.Todo
	_ = one.Decode(&todo)
	return todo
}

func (r TodoRepository) UpdateTask(i interface{}) interface{} {
	todo := i.(model.Todo)
	todo.Status = !todo.Status
	filter := bson.M{"_id": todo.ID}
	_, err := r.collection.UpdateOne(context.Background(), filter, bson.D{
		{"$set", bson.D{{"status", todo.Status}}},
	})

	if err != nil {
		return err
	}

	one := r.collection.FindOne(context.Background(), filter)
	_ = one.Decode(&todo)
	return todo
}

func (r TodoRepository) DeleteTask(id string) {
	//fmt.Println(objectId)
	filter := bson.M{"_id": id}
	_, _ = r.collection.DeleteOne(context.Background(), filter)
}

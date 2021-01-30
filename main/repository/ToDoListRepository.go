package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"todo/main/config"
	"todo/main/model/entity"
)


func SaveToDoList(t entity.ToDo) entity.ToDo {
	_, err := config.Database.Collection("todo").InsertOne(context.TODO(), t)

	if err != nil {
		panic(err)
	}

	return t
}

func GetAll() []entity.ToDo {

	var todos []entity.ToDo
	find, err := config.Database.Collection("todo").Find(context.TODO(), bson.D{})

	if err != nil {
		panic(err)
	}

	defer find.Close(context.TODO())

	for find.Next(context.TODO()) {
		var t entity.ToDo
		find.Decode(&t)
		todos = append(todos, t)
	}
	return todos
}
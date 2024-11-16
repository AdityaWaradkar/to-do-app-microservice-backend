package models

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var todoCollection *mongo.Collection

type Todo struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Completed   bool               `bson:"completed" json:"completed"`
	UserID      primitive.ObjectID `bson:"user_id" json:"user_id"`
}

type TodoInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	UserID      string `json:"user_id"`
}

// ConnectDB connects to MongoDB and sets up the todoCollection
func ConnectDB(uri string) error {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	// Establish a connection with timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		return err
	}

	// Ping to check if the MongoDB connection is alive
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}

	// Assign the collections to the global variables
	todoCollection = client.Database("to-do-list-app").Collection("todos")

	log.Println("Connected to MongoDB successfully")
	return nil
}

// Save inserts a new todo document into MongoDB
func (t *Todo) Save() error {
	_, err := todoCollection.InsertOne(context.Background(), t)
	return err
}

// UpdateTodo updates a todo document in MongoDB
func UpdateTodo(id string, input TodoInput) error {
	todoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.M{"$set": bson.M{
		"title":       input.Title,
		"description": input.Description,
		"completed":   input.Completed,
	}}

	_, err = todoCollection.UpdateOne(context.Background(), bson.M{"_id": todoID}, update)
	return err
}

// DeleteTodo deletes a todo document from MongoDB
func DeleteTodo(id string) error {
	todoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = todoCollection.DeleteOne(context.Background(), bson.M{"_id": todoID})
	return err
}

// FetchTodos fetches all todos for a given user
func FetchTodos(userID primitive.ObjectID) ([]Todo, error) {
	var todos []Todo

	cursor, err := todoCollection.Find(context.Background(), bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo Todo
		if err := cursor.Decode(&todo); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

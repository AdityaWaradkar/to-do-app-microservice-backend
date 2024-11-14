package models

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		return err
	}
	if err := client.Ping(ctx, nil); err != nil {
		return err
	}

	todoCollection = client.Database("to-do-list-app").Collection("todos")
	log.Println("Connected to MongoDB successfully")

	// Insert dummy tasks after establishing connection
	insertDummyTasks()

	return nil
}

// InsertDummyTasks inserts 2-3 dummy tasks into the database
func insertDummyTasks() {
	// Check if there are any existing tasks before inserting
	count, err := todoCollection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		log.Println("Error checking task count:", err)
		return
	}

	if count > 0 {
		log.Println("Dummy tasks already inserted.")
		return
	}

	todos := []Todo{
		{
			Title:       "Complete Go tutorial",
			Description: "Finish the Go tutorial to learn the basics of Go programming.",
			Completed:   false,
			UserID:      primitive.NewObjectID(), // Replace with a valid user ID
		},
		{
			Title:       "Buy groceries",
			Description: "Get groceries for the week, including fruits and vegetables.",
			Completed:   false,
			UserID:      primitive.NewObjectID(), // Replace with a valid user ID
		},
		{
			Title:       "Read a book",
			Description: "Read a chapter of a programming book to improve skills.",
			Completed:   true,
			UserID:      primitive.NewObjectID(), // Replace with a valid user ID
		},
	}

	// Insert dummy tasks into the collection
	for _, todo := range todos {
		_, err := todoCollection.InsertOne(context.Background(), todo)
		if err != nil {
			log.Println("Error inserting dummy task:", err)
		}
	}

	log.Println("Dummy tasks inserted successfully")
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

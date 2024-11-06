package models

import (
    "context"
    "errors"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson/primitive" 
    "log"
    "time"
)

var todoCollection *mongo.Collection

// ToDo represents the structure of a to-do item.
type ToDo struct {
    ID          string    `json:"id" bson:"_id,omitempty"`
    Title       string    `json:"title" bson:"title"`
    Description string    `json:"description" bson:"description"`
    Completed   bool      `json:"completed" bson:"completed"`
    CreatedAt   time.Time `json:"created_at" bson:"created_at"`
}

// ConnectDB establishes a connection to the MongoDB database.
func ConnectDB(uri string) error {
    client, err := mongo.NewClient(options.Client().ApplyURI(uri))
    if err != nil {
        return err
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        return err
    }

    log.Println("Connected to MongoDB!")

    todoCollection = client.Database("to-do-list-app").Collection("todos")
    return nil
}

// AddToDo adds a new to-do item to the database.
func AddToDo(todo ToDo) (ToDo, error) {
    todo.CreatedAt = time.Now()
    result, err := todoCollection.InsertOne(context.Background(), todo)
    if err != nil {
        return ToDo{}, err
    }

    todo.ID = result.InsertedID.(primitive.ObjectID).Hex() // Ensure correct ID conversion
    return todo, nil
}

// GetToDo retrieves a to-do item by its ID.
func GetToDoByID(id primitive.ObjectID) (ToDo, error) { // Change parameter to primitive.ObjectID
    var todo ToDo
    err := todoCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&todo)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return ToDo{}, errors.New("to-do not found")
        }
        return ToDo{}, err
    }
    return todo, nil
}

// UpdateToDo updates an existing to-do item.
func UpdateToDoByID(id primitive.ObjectID, updatedToDo ToDo) (ToDo, error) { // Change parameter to primitive.ObjectID
    _, err := todoCollection.UpdateOne(
        context.Background(),
        bson.M{"_id": id},
        bson.M{"$set": updatedToDo},
    )
    if err != nil {
        return ToDo{}, err
    }

    updatedToDo.ID = id.Hex() // Make sure to return the correct ID format
    return updatedToDo, nil
}

// DeleteToDo removes a to-do item from the database.
func DeleteToDoByID(id primitive.ObjectID) error { // Change parameter to primitive.ObjectID
    _, err := todoCollection.DeleteOne(context.Background(), bson.M{"_id": id})
    if err != nil {
        return err
    }
    return nil
}

// GetAllToDos retrieves all to-do items from the database.
func GetAllToDos() ([]ToDo, error) {
    cursor, err := todoCollection.Find(context.Background(), bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())

    var todos []ToDo
    for cursor.Next(context.Background()) {
        var todo ToDo
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

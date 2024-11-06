package models

import (
    "context"
    "log"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

var notificationCollection *mongo.Collection

// Notification represents the structure of a notification.
type Notification struct {
    ID        string    `json:"id" bson:"_id,omitempty"`
    Message   string    `json:"message" bson:"message"`
    CreatedAt time.Time `json:"created_at" bson:"created_at"`
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

    notificationCollection = client.Database("to-do-list-app").Collection("notifications")
    return nil
}

// AddNotification adds a new notification to the database.
func AddNotification(message string) (Notification, error) {
    notification := Notification{
        Message:   message,
        CreatedAt: time.Now(),
    }

    result, err := notificationCollection.InsertOne(context.Background(), notification)
    if err != nil {
        return Notification{}, err
    }

    notification.ID = result.InsertedID.(primitive.ObjectID).Hex() // Use primitive.ObjectID
    return notification, nil
}

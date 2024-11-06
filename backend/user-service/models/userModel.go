package models

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var client *mongo.Client
var userCollection *mongo.Collection

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"` // Use primitive.ObjectID
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
	Email    string             `json:"email" bson:"email"`
}

// ConnectDB establishes a connection to the MongoDB database.
func ConnectDB(uri string) error {
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	// Ping the database to verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	log.Println("Connected to MongoDB!")

	// Set the user collection
	userCollection = client.Database("to-do-list-app").Collection("users")

	return nil
}

// Register a new user
func RegisterUser(newUser User) (User, error) {
    // Log the new user data
    log.Printf("Attempting to register user: %+v\n", newUser)

    // Insert the new user into the database
    newUser.ID = primitive.NewObjectID() // Set the ID manually
    _, err := userCollection.InsertOne(context.Background(), newUser)
    if err != nil {
        log.Printf("Error inserting new user: %s\n", err) // Log the error
        return User{}, err
    }

    log.Println("User registered successfully")
    return newUser, nil
}



// Authenticate user
func AuthenticateUser(username, password string) (User, error) {
	user := User{}
	err := userCollection.FindOne(context.Background(), bson.M{"username": username, "password": password}).Decode(&user)
	if err != nil {
		return User{}, errors.New("invalid username or password")
	}
	return user, nil
}

// Update user profile
func UpdateUserProfile(id string, updatedUser User) (User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return User{}, errors.New("invalid user ID")
	}
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{"email": updatedUser.Email}}

	result := userCollection.FindOneAndUpdate(context.Background(), filter, update)
	if result.Err() != nil {
		return User{}, errors.New("user not found")
	}

	return updatedUser, nil
}

// Reset password
func ResetPassword(username, newPassword string) error {
	filter := bson.M{"username": username}
	update := bson.M{"$set": bson.M{"password": newPassword}} // Hash the password in production

	result := userCollection.FindOneAndUpdate(context.Background(), filter, update)
	if result.Err() != nil {
		return errors.New("user not found")
	}
	return nil
}

// Delete user account
func DeleteUser(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid user ID")
	}
	filter := bson.M{"_id": objID}
	result, err := userCollection.DeleteOne(context.Background(), filter)
	if err != nil || result.DeletedCount == 0 {
		return errors.New("user not found")
	}
	return nil
}

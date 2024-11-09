package models

import (
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
    "golang.org/x/crypto/bcrypt"
    "time"
)

var userCollection *mongo.Collection

type User struct {
    ID       primitive.ObjectID `bson:"_id,omitempty"`
    Username string             `bson:"username"`
    Email    string             `bson:"email"`
    Password string             `bson:"password"`
}

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
    if err := client.Ping(ctx, readpref.Primary()); err != nil {
        return err
    }

    userCollection = client.Database("to-do-list-app").Collection("users")
    return nil
}

func (u *User) Save() error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    u.Password = string(hashedPassword)
    _, err = userCollection.InsertOne(context.Background(), u)
    return err
}

func FindUserByEmail(email string) (*User, error) {
    var user User
    err := userCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (u *User) CheckPassword(password string) error {
    return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

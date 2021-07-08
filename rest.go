// Here Dealing with CRUD Rest Api along with MongoDB Database;
/*
	Details Steps are mentioned here:
	1. First of all let's install Mongo-DB Driver and import the libraries required for it.
	2. Creating a function for database configuration
	3.

*/
package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Now creating a function for Database Configuration

func connect() {
	// Configuring the database.
	clientOptions := options.Client().ApplyURI("mongodb_url")
	client, err := mongo.NewClient(clientOptions)

	// Setting up context required by the mongo.Connect()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	// canceling the context for avoiding memory leak
	defer cancel()

	// Ping our db connection
	err = client.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}
	// Connect to the database
	db := client.Database("go_mongo")
	TodoCollection(db)

	return
}

// Now database has been configured
// Now creating an instance of it

/*
	Structure for storing->
	1. ID of todo
	2. Title of todo
	3. Body of todo
	4. Time when todo is completed
	5. Time when todo is updated

*/

type Todo struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Completed string    `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

/*
	Following are the Todo's for->
	1. Getting all todo's
	2. Creating todo's
	3. Getting Single todo's
	4. Editing the todo's
	5. Deleting the todo's

*/

// DATABASE INSTANCE
var collection *mongo.Collection

func TodoCollection(c *mongo.Database) {
	collection = c.Collection("todos")
}

func GetAllTodos(c *gin.Context) {

}

func CreateTodo(c *gin.Context) {

}

func GetSingleTodo(c *gin.Context) {

}

func EditTodo(c *gin.Context) {

}

func DeleteTodo(c *gin.Context) {

}

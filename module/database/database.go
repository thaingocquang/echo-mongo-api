package database

import (
	"TestGoAPI/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var (
	db     *mongo.Database
	client *mongo.Client
)

// Connect ...
func Connect() {
	envVars := config.GetEnv()

	// Connect
	cl, err := mongo.NewClient(options.Client().ApplyURI(envVars.Database.URI))
	if err != nil {
		log.Println(err)
		log.Fatal("Cannot connect to database:", envVars.Database.URI)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = cl.Connect(ctx)
	if err != nil {
		log.Println(err)
	}
	client = cl
	db = cl.Database(envVars.Database.Name)
	fmt.Println("Database Connected to", envVars.Database.Name)
}

package database 

import (
	"context"
	"log"
  "fmt"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBConnection(){
  
  // Handle the loading of environement variables
  envVar, err := godotenv.Read("../.env")
  if err != nil {
    log.Fatal("Error loading .env file. Check if the file exists with proper permissions.")
  }

  // Setting up connection
  mongoURL := fmt.Sprint(
    envVar["MONGO_PROTOCOL"], 
    // envVar["MONGO_USER"], ":",
    // envVar["MONGO_PASSWORD"], "@",
    envVar["MONGO_HOSTNAME"], ":", 
    envVar["MONGO_PORT"], "/",
    )
  fmt.Println(mongoURL)

  serverAPI := options.ServerAPI(options.ServerAPIVersion1)
  opts := options.Client().ApplyURI(mongoURL).SetServerAPIOptions(serverAPI)
  mongoClient, err := mongo.Connect(context.TODO(), opts)
  if err != nil {
    panic(err)
  }
  defer func() {
    if err = mongoClient.Disconnect(context.TODO()); err != nil {
      panic(err)
    }
  }()

  // Checking connection
  var result bson.M
  if err := mongoClient.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
    panic(err)
  }
  fmt.Println("Pinged MongoDB. Connection successful.")

}


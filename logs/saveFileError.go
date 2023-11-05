package logs

import (
	"context"
	"fmt"
	"github.com/Fabricio2210/mongoDb"
	"github.com/joho/godotenv"
	"os"
	"time"
)

type ErrorObject struct {
	FileName  string
	Error     string
	Folder    string
	Service   string
	UpdatedAt time.Time
}

func SaveFileError(fileName string, folder string, service string, error string) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}
	if err := mongodb.ConnectDB(os.Getenv("URL")); err != nil {
		fmt.Println(err)
	}
	defer mongodb.DisconnectDB()
	collection := mongodb.Client.Database(os.Getenv("DB")).Collection("chatlogsErrors")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	errorObject := ErrorObject{FileName: fileName, Folder: folder, Service: service, Error: error, UpdatedAt: time.Now()}

	result, err := collection.InsertOne(ctx, errorObject)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Inserted document ID:", result.InsertedID)
}

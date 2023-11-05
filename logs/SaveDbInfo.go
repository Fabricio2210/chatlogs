package logs

import (
	"context"
	"fmt"
	"github.com/Fabricio2210/mongoDb"
	"github.com/joho/godotenv"
	"os"
	"time"
)

type DbInfoObject struct {
	FileList  []string
	Memory    uint64
	Folder    string
	ExecTime  string
	UpdatedAt time.Time
}

func SaveDbInfo(fileList []string, folder string, memory uint64, execTime string) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}
	if err := mongodb.ConnectDB(os.Getenv("URL")); err != nil {
		fmt.Println(err)
	}
	defer mongodb.DisconnectDB()
	collection := mongodb.Client.Database(os.Getenv("DB")).Collection("chatlogsDbInfo")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	dbInfoObject := DbInfoObject{FileList: fileList, Folder: folder, Memory: memory, ExecTime: execTime, UpdatedAt: time.Now()}
	if len(fileList) == 0 {
		fmt.Println("There is no data to save")
	} else {
		result, err := collection.InsertOne(ctx, dbInfoObject)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Inserted document ID:", result.InsertedID)
	}
}

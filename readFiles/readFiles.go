package readFiles

import (
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"os"
)

func ReadFiles(subject string) []string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}
	var fileList []string
	files, err := ioutil.ReadDir(fmt.Sprintf("%s/%s", os.Getenv("CHANNELS_PATH"), subject))
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			fmt.Printf("Directory: %s\n", file.Name())
		} else {
			fileList = append(fileList, file.Name())
		}
	}
	return fileList
}

package savedata

import (
	"bytes"
	"log"
	"net/http"
)

func Savedata(data []byte) error {
	req, err := http.Post("http://localhost:9200/_bulk", "application/json", bytes.NewReader(data))
	if err != nil {
		log.Println("Error creating POST request:", err)
		return nil
	}
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Println(err)
	}
	return nil
}

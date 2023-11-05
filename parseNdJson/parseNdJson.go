package parseNdJson

import (
	"encoding/json"
	"fmt"
	changedateformat "github.com/Fabricio2210/dateFormat"
	"github.com/Fabricio2210/elastic"
	"github.com/Fabricio2210/logs"
	parsestring "github.com/Fabricio2210/parseString"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

func ParseNdJson(fileName string, subject string) []byte {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}
	var sliceDoc []elastic.MyDocument
	data, err := os.ReadFile(fmt.Sprintf("%s/%s/%s", os.Getenv("CHANNELS_PATH"), subject, fileName))
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		if line == "" {
			continue
		}
		data, err := parsestring.Parsestring(line)
		if err != nil {
			errorMessage := fmt.Sprintf("error %s at the line %d from file %s from folder %s.", err, i+1, fileName, subject)
			logs.SaveFileError(fileName, subject, os.Getenv("SERVICE"), errorMessage)
			fmt.Println(errorMessage)
			continue
		}
		parsedDate, err := changedateformat.ChangeDateFormat(data.Date)
		if err != nil {
			errorMessage := fmt.Sprintf("error %s at the line %d from file %s from folder %s.", err, i+1, fileName, subject)
			logs.SaveFileError(fileName, subject, os.Getenv("SERVICE"), errorMessage)
			fmt.Println(errorMessage)
			continue
		}
		doc := elastic.MyDocument{
			UserName: data.Name,
			Text:     data.Message,
			Hour:     data.Hour,
			LogDay:   data.Date,
			Date:     parsedDate,
			Subject:  subject,
		}
		sliceDoc = append(sliceDoc, doc)

	}
	var ndjsonBytes []byte
	for _, doc := range sliceDoc {
		indexObject := map[string]interface{}{
			"index": map[string]string{
				"_index": os.Getenv("SCHEMA"),
			},
		}
		indexJSON, err := json.Marshal(indexObject)
		if err != nil {
			fmt.Println("Error marshaling index object to JSON:", err)
		}
		ndjsonBytes = append(ndjsonBytes, indexJSON...)
		ndjsonBytes = append(ndjsonBytes, []byte("\n")...)

		docJSON, err := json.Marshal(doc)
		if err != nil {
			fmt.Println("Error marshaling document to JSON:", err)
		}
		ndjsonBytes = append(ndjsonBytes, docJSON...)
		ndjsonBytes = append(ndjsonBytes, []byte("\n")...)
	}
	return ndjsonBytes
}

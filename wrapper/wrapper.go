package wrapper

import (
	"fmt"
	"github.com/Fabricio2210/logs"
	"github.com/Fabricio2210/parseNdJson"
	"github.com/Fabricio2210/readFiles"
	"github.com/Fabricio2210/saveData"
	"runtime"
	"time"
)

func Wrapper(subject string) {
	start := time.Now()
	var m runtime.MemStats
	fileList := readFiles.ReadFiles(subject)
	for _, fileName := range fileList {
		ndJson := parseNdJson.ParseNdJson(fileName, subject)
		savedata.Savedata((ndJson))
	}
	elapsed := time.Since(start)
	runtime.ReadMemStats(&m)
	memory := m.Alloc / 1024 / 1024
	timeParsed := fmt.Sprintf("%v", elapsed)
	fmt.Printf("Saving folder %s \n", subject)
	fmt.Printf("Alloc = %v MiB \n", memory)
	fmt.Printf("function took %s \n", elapsed)
	logs.SaveDbInfo(fileList, subject, memory, timeParsed)
}

package main

import (
	"CrawlFileSystem/models"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	path2 "path"
	"path/filepath"
	"sync"
	"time"
)

var tasks sync.WaitGroup
var GLOBALARGUMENTS = &models.CommandLineArgument{}

func main() {
	println("Started")

	flag.StringVar(&GLOBALARGUMENTS.BaseSystemDirector, "base", "", "help message for flagname")
	flag.StringVar(&GLOBALARGUMENTS.ApiUrl, "url", "", "Missing API Url")
	GLOBALARGUMENTS.TimeBetweenRequest = *flag.Duration("timeBetweenRequest", 20, "Missing Time Between Request")

	flag.Parse()
	iterate(GLOBALARGUMENTS.BaseSystemDirector)
	tasks.Wait()

	println("Complete")
}

func iterate(path string) {
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return filepath.SkipDir
		}

		var request = *models.NewFileRequest(info.Name(), path, path2.Ext(info.Name()))
		tasks.Add(1)
		go func(r models.FilesRequest) {
			defer tasks.Done()
			LogFile(r)
		}(request)

		time.Sleep(GLOBALARGUMENTS.TimeBetweenRequest * time.Millisecond)
		return nil
	})
}

func LogFile(request models.FilesRequest) {

	method := "POST"

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(request)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, GLOBALARGUMENTS.ApiUrl, &buf)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
}

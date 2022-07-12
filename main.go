package main

import (
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

var wg sync.WaitGroup

func main() {
	baseDirector := ""

	flag.StringVar(&baseDirector, "base", "jk", "help message for flagname")
	flag.Parse()
	iterate(baseDirector)
	wg.Wait()
}

func iterate(path string) {
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return filepath.SkipDir
			//log.Fatalf(err.Error())

		}

		var request = FilesRequest{
			Name:        info.Name(),
			Location:    path2.Join(path, info.Name()),
			Description: "",
			Extension:   path2.Ext(info.Name()),
			CreatedOn:   time.Now(),
		}
		wg.Add(1)
		go func(r FilesRequest) {
			// Decrement the counter when the go routine completes
			defer wg.Done()
			// Call the function check
			LogFile(r)
		}(request)

		fmt.Printf("File Name: %s\n", info.Name())
		time.Sleep(20 * time.Millisecond)
		return nil
	})
}

func LogFile(request FilesRequest) {
	url := "https://localhost:7124/Files/"
	method := "POST"

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(request)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, &buf)

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

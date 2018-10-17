package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func isOlder(t time.Time, days int) bool {
	return time.Now().Sub(t) > 1*time.Hour
}

func getFiles(location string, days int) ([]string, error) {
	var files []string
	err := filepath.Walk(location, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	test := make([]string, 0)

	for _, file := range files {
		at, err := os.Stat(file)
		if err != nil {
			log.Fatal(err)
		}
		if isOlder(at.ModTime(), days) {
			test = append(test, file)
		}

	}

	return test, err

}

func main() {
	var location string
	var days int
	flag.StringVar(&location, "location", "", "location to scan")
	flag.IntVar(&days, "days", 0, "days to poll")
	flag.Parse()

	files, err := getFiles(location, days)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file)
	}

}

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
	hours := time.Duration(days*24) * time.Hour
	return time.Now().Sub(t) > hours*time.Hour
}

func getFiles(location string, days int) error {
	var files []string
	err := filepath.Walk(location, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file)
	}

	return err

}

func main() {
	var location string
	var days int
	flag.StringVar(&location, "location", "", "location to scan")
	flag.IntVar(&days, "days", 0, "days to poll")
	flag.Parse()

	_ = getFiles(location, days)

}

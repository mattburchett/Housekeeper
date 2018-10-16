package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	var location string
	var days int
	flag.StringVar(&location, "location", "", "location to scan")
	flag.IntVar(&days, "days", 0, "days to poll")
	flag.Parse()

	var files []string
	err := filepath.Walk(location, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now()
	cutoff := now.Add(-int(days * time.Hour)

	fmt.Println(cutoff)

	for _, file := range files {
		if diff := now.Sub(files.ModTime()); diff > cutoff {
			fmt.Println(diff)
		}
	}

}

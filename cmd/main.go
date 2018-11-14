package main

import (
	"flag"
	"log"

	"git.linuxrocker.com/mattburchett/Housekeeper/pkg/config"
	"git.linuxrocker.com/mattburchett/Housekeeper/pkg/locator"
)

// func getFiles(location string, days int) ([]string, error) {
// 	var files []string
// 	err := filepath.Walk(location, func(path string, info os.FileInfo, err error) error {
// 		files = append(files, path)
// 		return nil
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	test := make([]string, 0)

// 	for _, file := range files {
// 		at, err := os.Stat(file)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		if isOlder(at.ModTime(), days) {
// 			test = append(test, file)
// 		}

// 	}

// 	return test, err

// }

func main() {
	var c string
	var days int
	var sectionID int

	flag.StringVar(&c, "c", "", "Configuration to load")
	flag.IntVar(&days, "days", 0, "days to poll")
	flag.IntVar(&sectionID, "sectionid", 0, "pick a section ID")
	flag.Parse()
	if c == "" {
		log.Fatal("You need to specify a configuration file.")
	}
	if sectionID == 0 {
		log.Fatal("You need to specify a section ID for Plex.")
	}

	cfg, err := config.GetConfig(c)
	if err != nil {
		log.Fatal(err)
	}
	locator.GetCount(cfg, sectionID)
}

package main

import (
	"flag"
	"log"

	"git.linuxrocker.com/mattburchett/Housekeeper/pkg/communicator"
	"git.linuxrocker.com/mattburchett/Housekeeper/pkg/config"
	"git.linuxrocker.com/mattburchett/Housekeeper/pkg/eraser"
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
	var check bool
	var delete bool

	flag.StringVar(&c, "c", "", "Configuration to load")
	flag.IntVar(&days, "days", 0, "days to poll")
	flag.IntVar(&sectionID, "sectionid", 0, "pick a section ID")
	flag.BoolVar(&check, "check", true, "Perform only a check. Do not delete.")
	flag.BoolVar(&delete, "delete", false, "Perform the delete task.")
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

	ids, titles := locator.GetTitles(cfg, sectionID, days)

	if check {
		err = communicator.TelegramPost(cfg, titles)
		if err != nil {
			log.Fatal(err)
		}
	}

	if delete {
		files := eraser.LookupFileLocation(cfg, ids)
		err = eraser.DeleteMedia(delete, files)
		if err != nil {
			log.Println(err)
		}
	}

}

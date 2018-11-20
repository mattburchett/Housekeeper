package main

import (
	"flag"
	"fmt"
	"log"

	"git.linuxrocker.com/mattburchett/Housekeeper/pkg/communicator"
	"git.linuxrocker.com/mattburchett/Housekeeper/pkg/config"
	"git.linuxrocker.com/mattburchett/Housekeeper/pkg/eraser"
	"git.linuxrocker.com/mattburchett/Housekeeper/pkg/locator"
)

func main() {
	var c string
	var days int
	var sectionID int
	var check bool
	var delete bool

	flag.StringVar(&c, "config", "", "Configuration to load")
	flag.IntVar(&days, "days", 0, "How many days of inactivity to look for on Plex.")
	flag.IntVar(&sectionID, "sectionid", 0, "Plex Section ID")
	flag.BoolVar(&check, "check", true, "Perform only a check. This will send the message out to Telegram with what can be removed. Does not delete.")
	flag.BoolVar(&delete, "delete", false, "Perform the delete task.")
	flag.Parse()

	// Stop the app if they're missing required flags.
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

	libraryType := locator.GetLibraryType(cfg, sectionID)

	ids, titles := locator.GetTitles(cfg, sectionID, days)

	if check {
		err = communicator.TelegramPost(cfg, titles)
		if err != nil {
			log.Fatal(err)
		}
	}

	files := eraser.LookupTVFileLocation(cfg, ids)
	fmt.Printf("%v\n", files)

	if delete {
		if libraryType == "movie" {
			files := eraser.LookupMovieFileLocation(cfg, ids)
			err = eraser.DeleteMovies(delete, files)
			if err != nil {
				log.Println(err)
			}
		} else if libraryType == "show" {
		}
	}
}

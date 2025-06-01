package main

import (
	"log"
	flagger "scrappyscrap/internal/Flagger"
	scrapper "scrappyscrap/internal/Scrapper"
)

func main() {
	flags := flagger.FlaggerSetup().Do()
	Scrapper := scrapper.NewScrapper(flags)
	switch len(flags) {
	case 1:
		if err := Scrapper.ScrapPage(); err != nil {
			log.Fatalf("%v", err)
		}
	default:
		if err := Scrapper.ScrapPages(); err != nil {
			log.Fatalf("%v", err)
		}

	}
}

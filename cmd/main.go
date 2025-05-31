package main

import (
	"fmt"
	flagger "scrappyscrap/internal/Flagger"
	scrapper "scrappyscrap/internal/Scrapper"
)

func main() {
	flags := flagger.FlaggerSetup().Do()
	Scrapper := scrapper.NewScrapper(flags)
	var neededElement string
	fmt.Println("Enter element what you want to parse: ")
	fmt.Scan(&neededElement)
	switch len(flags) {
	case 1:
		Scrapper.ScrapPage(neededElement)
	default:
		Scrapper.ScrapPages(neededElement)

	}
}

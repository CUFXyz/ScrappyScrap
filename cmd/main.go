package main

import (
	flagger "scrappyscrap/internal/Flagger"
	scrapper "scrappyscrap/internal/Scrapper"
)

func main() {
	flags := flagger.FlaggerSetup().Do()
	Scrapper := scrapper.NewScrapper(flags)
	Scrapper.ScrapPage("p")
}

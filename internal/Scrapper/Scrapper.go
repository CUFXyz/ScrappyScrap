package scrapper

import (
	"log"
	"os"
	"strings"

	"github.com/go-rod/rod"
)

type Scrapper struct {
	links []string
}

// Declaring Scrapper instance
func NewScrapper(links []string) *Scrapper {
	return &Scrapper{
		links: links,
	}
}

// Scrapping only one page, put an needed element into arguments
// Making txt file with parsed text with element from argument
func (scrap *Scrapper) ScrapPage(neededElement string) {
	var data []string
	page := rod.New().NoDefaultDevice().MustConnect().MustPage(scrap.links[0]).MustWaitLoad()
	defer page.MustClose()
	paragraphs := page.MustElements(neededElement)
	for _, elem := range paragraphs {
		data = append(data, elem.MustText())
	}
	formatteddata := strings.Join(data, " ")
	file, err := os.Create("Result.txt")
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer file.Close()
	file.Write([]byte(formatteddata))
}

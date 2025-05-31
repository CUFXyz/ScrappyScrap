package scrapper

import (
	"fmt"
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
func (scrap *Scrapper) ScrapPage(neededElement string) error {
	var data []string
	page := rod.New().NoDefaultDevice().MustConnect().MustPage(scrap.links[0]).MustWaitLoad()
	defer page.MustClose()
	elements := page.MustElements(neededElement)
	for _, elem := range elements {
		data = append(data, elem.MustText())
	}
	formatteddata := strings.Join(data, "\n")
	file, err := os.Create("Result.txt")
	if err != nil {
		return fmt.Errorf("fail at creating file")
	}
	defer file.Close()
	file.Write([]byte(formatteddata))
	return nil
}

// Scrapping multiple pages, also can scrap only one
// That will be pretty stupid to return an error because more than 1 link given
// By the way it also make txt file with parsed text from argument
// It will separate result from parsing links with [Link %number of link%]
func (scrap *Scrapper) ScrapPages(neededElement string) error {
	var data []string

	file, err := os.Create("MultipleResult.txt")
	if err != nil {
		return fmt.Errorf("fail at creating file")
	}
	defer file.Close()

	for i := range scrap.links {
		page := rod.New().NoDefaultDevice().MustConnect().MustPage(scrap.links[i]).MustWaitLoad()
		elements := page.MustElements(neededElement)
		data = append(data, fmt.Sprintf("\n[Link %d]\n", i+1))
		for _, elem := range elements {
			data = append(data, elem.MustText())
		}
		page.Close()
	}

	formatteddata := strings.Join(data, "\n")
	file.Write([]byte(formatteddata))
	return nil
}

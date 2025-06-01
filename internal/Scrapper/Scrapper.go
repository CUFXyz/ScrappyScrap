package scrapper

import (
	"fmt"
	"os"
	"strings"
	"sync"

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

// Scrapping multiple links, also can scrap only one, but if you want to parse only one link, use ScrapPage(), that's a lot faster than this method
// That will be pretty stupid to return an error because more than 1 link given
// By the way it also make txt file with parsed text from argument
// It will separate result from parsing links with [Link %number of link%]
func (scrap *Scrapper) ScrapPages(neededElement string) error {
	var data []string
	var wg sync.WaitGroup

	file, err := os.Create("MultipleResult.txt")
	if err != nil {
		return fmt.Errorf("fail at creating file")
	}
	defer file.Close()

	wg.Add(len(scrap.links))
	for i := range scrap.links {
		go func(i int) {
			defer wg.Done()
			var newdata []string
			page := rod.New().NoDefaultDevice().MustConnect().MustPage(scrap.links[i]).MustWaitLoad()
			defer page.Close()
			elements := page.MustElements(neededElement)
			newdata = append(newdata, fmt.Sprintf("\n[Link %d]\n", i+1))
			for _, elem := range elements {
				newdata = append(newdata, elem.MustText())
			}
			result := strings.Join(newdata, " ")
			data = append(data, result)
		}(i)
	}
	wg.Wait()
	formatteddata := strings.Join(data, "\n")
	file.Write([]byte(formatteddata))
	return nil
}

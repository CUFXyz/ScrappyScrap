package scrapper

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/go-rod/rod"
)

type Scrapper struct {
	pages   []*rod.Page
	element string
}

// Creating an array of pages for scrapper for work
func SetupPages(links []string) []*rod.Page {
	var pages []*rod.Page
	for i := range links {
		p := rod.New().NoDefaultDevice().MustConnect().MustPage(links[i]).MustWaitLoad()
		pages = append(pages, p)
	}
	return pages
}

// Constructor for Scrapper instance
func NewScrapper(links []string) *Scrapper {
	var neededElement string
	fmt.Println("Enter element what you want to parse: ")
	fmt.Scan(&neededElement)
	return &Scrapper{
		pages:   SetupPages(links),
		element: neededElement,
	}
}

// Scrapping only one page, put an needed element into arguments
// Making txt file with parsed text with element from argument
func (scrap *Scrapper) ScrapPage() error {
	var data []string
	defer scrap.pages[0].MustClose()

	file, err := os.Create("SingleResult.txt")
	if err != nil {
		return fmt.Errorf("fail at creating file")
	}
	defer file.Close()

	elements := scrap.pages[0].MustElements(scrap.element)
	for _, element := range elements {
		data = append(data, element.MustText())
	}

	result := strings.Join(data, " ")
	file.WriteString(result)
	return nil
}

// Scrapping multiple links, also can scrap only one, but if you want to parse only one link, use ScrapPage(), that's a lot faster than this method
// That will be pretty stupid to return an error because more than 1 link given
// By the way it also make txt file with parsed text from argument
// It will separate result from parsing links with [Link %number of link%]
func (scrap *Scrapper) ScrapPages() error {
	var data []string
	var wg sync.WaitGroup

	file, err := os.Create("MultipleResult.txt")
	if err != nil {
		return fmt.Errorf("fail at creating file")
	}
	defer file.Close()

	wg.Add(len(scrap.pages))
	for i := range scrap.pages {
		go func(i int) {
			defer wg.Done()
			var newdata []string
			defer scrap.pages[i].MustClose()
			elements := scrap.pages[i].MustElements(scrap.element)
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
	file.WriteString(formatteddata)
	return nil
}

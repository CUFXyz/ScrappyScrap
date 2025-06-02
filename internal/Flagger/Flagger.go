package flagger

import (
	"fmt"
	"os"
)

// Const flags for the flagger unit
const (
	Multiple string = "-m"
	Help     string = "-help"
)

type Flagger struct {
	flags []string
}

// Creating a instance for flag reader.
// Usage: go run main.go <special flag> <link> <link> (if -m flag given)
func FlaggerSetup() *Flagger {
	flags := os.Args[1:]
	return &Flagger{
		flags: flags,
	}
}

// Checking flags
// If no flags given - it's working with only one link
// If "-m" flag given - working with multiple links
func (f *Flagger) Do() []string {
	var newFlags []string
	if f.flags[0] == Multiple {
		newFlags = append(newFlags, f.flags[1:]...)
	}

	if f.flags[0] != Multiple {
		newFlags = append(newFlags, f.flags[0])
	}

	if f.flags[0] == Help {
		f.Help()
	}

	return newFlags
}

func (f *Flagger) Help() {
	fmt.Println("Usage:")
	fmt.Println("./ScrappyScrap -help               | Help information table")
	fmt.Println("./ScrappyScrap -m <link> <link>... | For multiple link scraping. Actually, you can scrap a single link with that flag, but not recommended, it's gonna take more time")
	fmt.Println("./ScrappyScrap <link>              | For single link scraping")
	os.Exit(0)
}

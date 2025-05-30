package flagger

import "os"

// Const flags for the flagger unit
const (
	Multiple string = "-m"
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

	return newFlags
}

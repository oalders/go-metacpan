package main

// go run main.go -q Mojo

import (
	"flag"
	"fmt"

	"github.com/oalders/go-metacpan"
)

func main() {
	var query string
	flag.StringVar(&query, "q", "", "search query")
	flag.Parse()

	dists, err := metacpan.SearchAutocomplete(query)

	if err != nil {
		fmt.Println(err)
		return
	}

	for i, dist := range dists {
		fmt.Println("---")
		fmt.Printf("%d\n", i+1)
		fmt.Printf("Name\t: %s\n", dist.Fields.Documentation)
		fmt.Printf("Author\t: %s\n", dist.Fields.Author)
		fmt.Printf("Release\t: %s\n", dist.Fields.Release)
		fmt.Printf("url\t: %s\n", dist.URL())
	}
}

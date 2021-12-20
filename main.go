package main

import (
	"flag"
	"log"

	"github.com/byron1st/dr-extractor-golang/lib"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if err := lib.ExtractCallgraph(args[0]); err != nil {
		log.Fatal(err)
	}
}

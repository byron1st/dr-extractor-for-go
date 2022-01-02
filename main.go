package main

import (
	"flag"
	"log"

	"github.com/byron1st/dr-extractor-golang/lib"
)

func main() {
	flag.Parse()
	args := flag.Args()
	pkgName := args[0]
	baseName := ""
	if len(args) > 1 {
		baseName = args[1]
	}
	if err := lib.ExtractCallgraph(pkgName, baseName); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"flag"
	"log"

	"github.com/jsquiroz/howmuch/pkg"
)

func main() {
	asset := flag.String("asset", "", "Cryptocurrencie name")

	flag.Parse()

	if *asset == "" {
		log.Fatal("Asset is empty")
	}
	pkg.Draw(*asset)
}

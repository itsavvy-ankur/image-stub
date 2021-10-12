package main

import (
	"flag"
	"log"
	"net/http"

	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//    sw "github.com/myname/myrepo/go"
	//
	sw "github.com/itsavvy-ankur/image-stub/go"
)

func main() {
	imgYamlConfigPath := flag.String("imgconfig", "<update_me>", "A config containing images to load")
	flag.Parse()
	log.Printf("Server started")

	router := sw.NewRouter(*imgYamlConfigPath)

	log.Fatal(http.ListenAndServe(":8080", router))
}

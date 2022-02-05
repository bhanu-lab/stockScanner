package main

import (
	"log"
	"os"
	"stockScanner/Scrapper"
	"stockScanner/filters"
)

/*
 * command line application for generating bullish/bearish stocks
 */
func main() {
	// Read command line arguments
	argsWithoutProg := os.Args[1:]
	allOptions := filters.GetCommandLineOptions(argsWithoutProg)
	log.Printf("All options received : %+v \n", allOptions)
	if ok := filters.ValidateCommandLineOptions(allOptions); !ok {
		log.Printf("validation failed please revisit options provided")
	} else {
		//log.Printf("all arguments %#v \n", argsWithProg)
		//log.Printf("arguments with program name %#v \n", argsWithoutProg)
		Scrapper.ScrapeContent(argsWithoutProg, allOptions)
	}
}

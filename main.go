package main

import (
	"log"
	"os"
	"stockScanner/Scrapper"
	"stockScanner/utils"
)

/*
 * command line application for generating bullish/bearish stocks
 */
func main() {
	// Read command line arguments
	argsWithoutProg := os.Args[1:]
	allOptions := utils.GetCommandLineOptions(argsWithoutProg)
	if ok := utils.ValidateCommandLineOptions(allOptions); !ok {
		log.Printf("validation failed ")
	} else {
		//log.Printf("all arguments %#v \n", argsWithProg)
		//log.Printf("arguments with program name %#v \n", argsWithoutProg)
		Scrapper.ScrapeContent(argsWithoutProg)
	}
}

package main

import (
	"os"
	"stockScanner/Scrapper"
)

func main() {
	//argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]
	//log.Printf("all arguments %#v \n", argsWithProg)
	//log.Printf("arguments with program name %#v \n", argsWithoutProg)
	Scrapper.ScrapeContent(argsWithoutProg)
}

// Scrapper provides scrapping API for scrapping stock related data with bullish and bearish stocks
package Scrapper

import (
	"log"
	"net/http"
	"net/url"
	"stockScanner/fileio"
	"stockScanner/requests"
	"stockScanner/types"
	"stockScanner/utils"
	"strings"

	"golang.org/x/net/html"
)

/*
 * ScrapeContent takes command line arguments for options
 * -t stock option type
 * -s simple technical details only
 */
func ScrapeContent(args []string) {
	index := 0
	// Create HTTP client with timeout
	client := &http.Client{}

	bearishStocksScanner := `{"action": "advanced_search", "info": {"0": {"cols": "last_close", "opts": "<", "cols1": "ema_20"}, "1": {"cols": "last_close", "opts": "<", "cols1": "sma_50"}, "2": {"cols": "dmi_plus", "opts": "<", "cols1": "dmi_minus"}, "3": {"cols": "", "opts": " like ", "cols1": ""}, "4": {"cols": "", "opts": " like ", "cols1": ""}, "5": {"cols": "avg_volume", "opts": ">", "strs": "500000"}, "6": {"cols": "atr", "opts": ">", "strs": "10"}, "7": {"cols": "adx", "opts": ">", "strs": "25"}, "8": {"cols": "rsi", "opts": ">", "strs": "30"}, "9": {"cols": "p_symbol", "opts": " not like ", "strs": "%-%"}}}`
	bullishStocksScanner := `{"action": "advanced_search", "info": {"0": {"cols": "last_close", "opts": ">", "cols1": "ema_5"}, "1": {"cols": "last_close", "opts": ">", "cols1": "ema_20"}, "2": {"cols": "last_close", "opts": "<", "cols1": "bband_upper"}, "3": {"cols": "", "opts": " like ", "cols1": ""}, "4": {"cols": "", "opts": " like ", "cols1": ""}, "5": {"cols": "atr", "opts": ">", "strs": "10"}, "6": {"cols": "adx", "opts": ">", "strs": "25"}, "7": {"cols": "avg_volume", "opts": ">", "strs": "500000"}, "8": {"cols": "p_symbol", "opts": " not like ", "strs": "%-%"}, "9": {"cols": "", "opts": " like ", "strs": ""}}}`
	params := url.Values{}
	stockScannerType := utils.GetStockScannerType(args)
	log.Printf("stock scanner type received is %s \n", stockScannerType)

	// sets required stockScannerType in params for URL encoding from command line argument
	if types.BULLISH == stockScannerType {
		log.Printf("creating bullish stock scanner object")
		params.Set("json", bullishStocksScanner)
	} else {
		log.Printf("creating bearish stock scanner object")
		params.Set("json", bearishStocksScanner)
	}
	value := params.Encode()
	/*log.Println("received encoded string is : ", value)*/

	err, htmlOutput := requests.CreateAPIRequestAndGetResponse(value, client)
	if err != nil {
		log.Panic("error while creating API request or while reading response")
	}
	//log.Println("****************************************************")
	//log.Printf("response : %+v \n", htmlOutput)
	log.Println("****************************************************")

	htmlTokens := html.NewTokenizer(strings.NewReader(htmlOutput))

	// outputs bullish or bearish based on the stockScannerType
	if types.BULLISH == stockScannerType {
		log.Println("*******************BULLISH STOCKS*******************")
	} else {
		log.Println("*******************BEARISH STOCKS*******************")
	}
	log.Println("****************************************************")

	var stocksData [][]string // holds data for both rows and columns when written to csv file
	var rowData []string      // holds data for each row
	i := 0
	for i < 1 {

		tt := htmlTokens.Next()
		//fmt.Printf("%T", tt)
		switch tt {
		case html.ErrorToken:
			//fmt.Println("End")
			i++
		case html.TextToken:
			//fmt.Println(tt.String())
		case html.StartTagToken:

			t := htmlTokens.Token()

			// if a new row tag is found enter and starting reading each column tag present
			isRow := t.Data == "tr"
			if isRow {
				//log.Printf("\n \n")
				if len(rowData) > 0 {

					stocksData = append(stocksData, rowData) //appending new row data to double array
					rowData = make([]string, 0)
				} else {
					log.Println("I am in else condition of len(rowData)")
					rowData = make([]string, 0)
				}

				_ = htmlTokens.Next() // navigate to next html token
				td := htmlTokens.Token()
				isAnchor := td.Data == "td" // check if current html tag is td

				if isAnchor { // if it is <td> tag then go ahead and extract data present inside <td>
					_ = htmlTokens.Next()
					eg := htmlTokens.Token()
					// getting data from <td> tag
					//log.Printf(strconv.Itoa(index) + "." +strings.ReplaceAll(eg.String(), "&lt;\\/td&gt;", ""))
					actualString := utils.ReplaceUnnecessaryHtmlData(eg)
					if actualString != "" {
						rowData = append(rowData, actualString)
					}
				}
				index++
				//for each column tag td present ger data present inside td tag and add it to th e string slice
			} else if t.Data == "td" {
				_ = htmlTokens.Next()
				eg := htmlTokens.Token()

				//log.Printf(" | "+strings.ReplaceAll(eg.String(), "&lt;\\/td&gt;", ""))
				actualString := utils.ReplaceUnnecessaryHtmlData(eg)
				if actualString != "" {
					rowData = append(rowData, actualString)
				}
			}
		}
	}

	csvFile, err := fileio.CreateCSVFile(stockScannerType)
	if err != nil {
		log.Panic("error while creating csv file check for errors while creating")
	}
	err = fileio.WriteCSVFile(csvFile, stocksData)
	if err != nil {
		log.Panic("not able to write to csv file error while writing to csv file")
	}
}

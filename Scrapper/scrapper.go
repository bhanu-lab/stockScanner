package Scrapper

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"stockScanner/fileio"
	"stockScanner/utils"
	"strings"

	"golang.org/x/net/html"
)

func ScrapeContent() {
	index := 0
	// Create HTTP client with timeout
	client := &http.Client{}

	//requestBytes := []byte(`{"action": "advanced_search", "info": {"0": {"cols": "last_close", "opts": ">", "cols1": "ema_5"}, "1": {"cols": "last_close", "opts": ">", "cols1": "ema_20"}, "2": {"cols": "last_close", "opts": "<", "cols1": "bband_upper"}, "3": {"cols": "", "opts": " like ", "cols1": ""}, "4": {"cols": "", "opts": " like ", "cols1": ""}, "5": {"cols": "atr", "opts": ">", "strs": "10"}, "6": {"cols": "adx", "opts": ">", "strs": "25"}, "7": {"cols": "avg_volume", "opts": ">", "strs": "500000"}, "8": {"cols": "p_symbol", "opts": " not like ", "strs": "%-%"}, "9": {"cols": "", "opts": " like ", "strs": ""}}}`)
	//requestBytes := []byte(`json={"action": "advanced_search", "info": {"0": {"cols": "last_close", "opts": ">", "cols1": "ema_5"}, "1": {"cols": "last_close", "opts": ">", "cols1": "ema_20"}, "2": {"cols": "last_close", "opts": "<", "cols1": "bband_upper"}, "3": {"cols": "", "opts": " like ", "cols1": ""}, "4": {"cols": "", "opts": " like ", "cols1": ""}, "5": {"cols": "atr", "opts": ">", "strs": "10"}, "6": {"cols": "adx", "opts": ">", "strs": "25"}, "7": {"cols": "avg_volume", "opts": ">", "strs": "500000"}, "8": {"cols": "p_symbol", "opts": " not like ", "strs": "%-%"}, "9": {"cols": "", "opts": " like ", "strs": ""}}}&_=`)

	//requestBytes := {"action": "advanced_search", "info": "{\"0\": {\"cols\": \"last_close\", \"opts\": \">\", \"cols1\": \"ema_5\"}, \"1\": {\"cols\": \"last_close\", \"opts\": \">\", \"cols1\": \"ema_20\"}, \"2\": {\"cols\": \"last_close\", \"opts\": \"<\", \"cols1\": \"bband_upper\"}, \"3\": {\"cols\": \"\", \"opts\": \" like \", \"cols1\": \"\"}, \"4\": {\"cols\": \"\", \"opts\": \" like \", \"cols1\": \"\"}, \"5\": {\"cols\": \"atr\", \"opts\": \">\", \"strs\": \"10\"}, \"6\": {\"cols\": \"adx\", \"opts\": \">\", \"strs\": \"25\"}, \"7\": {\"cols\": \"avg_volume\", \"opts\": \">\", \"strs\": \"500000\"}, \"8\": {\"cols\": \"p_symbol\", \"opts\": \" not like \", \"strs\": \"%-%\"}, \"9\": {\"cols\": \"\", \"opts\": \" like \", \"strs\": \"\"}}\\"}
	//value, err := json.Marshal(requestBytes)

	//var data = strings.NewReader(urlValues.Encode())
	params := url.Values{}
	params.Set("json", `{"action": "advanced_search", "info": {"0": {"cols": "last_close", "opts": ">", "cols1": "ema_5"}, "1": {"cols": "last_close", "opts": ">", "cols1": "ema_20"}, "2": {"cols": "last_close", "opts": "<", "cols1": "bband_upper"}, "3": {"cols": "", "opts": " like ", "cols1": ""}, "4": {"cols": "", "opts": " like ", "cols1": ""}, "5": {"cols": "atr", "opts": ">", "strs": "10"}, "6": {"cols": "adx", "opts": ">", "strs": "25"}, "7": {"cols": "avg_volume", "opts": ">", "strs": "500000"}, "8": {"cols": "p_symbol", "opts": " not like ", "strs": "%-%"}, "9": {"cols": "", "opts": " like ", "strs": ""}}}`)
	value := params.Encode()
	fmt.Println("received encoded string is : ", value)

	//var data = strings.NewReader(`json=%7B%22action%22%3A%20%22advanced_search%22%2C%20%22info%22%3A%20%7B%220%22%3A%20%7B%22cols%22%3A%20%22last_close%22%2C%20%22opts%22%3A%20%22%3E%22%2C%20%22cols1%22%3A%20%22ema_5%22%7D%2C%20%221%22%3A%20%7B%22cols%22%3A%20%22last_close%22%2C%20%22opts%22%3A%20%22%3E%22%2C%20%22cols1%22%3A%20%22ema_20%22%7D%2C%20%222%22%3A%20%7B%22cols%22%3A%20%22last_close%22%2C%20%22opts%22%3A%20%22%3C%22%2C%20%22cols1%22%3A%20%22bband_upper%22%7D%2C%20%223%22%3A%20%7B%22cols%22%3A%20%22%22%2C%20%22opts%22%3A%20%22%20like%20%22%2C%20%22cols1%22%3A%20%22%22%7D%2C%20%224%22%3A%20%7B%22cols%22%3A%20%22%22%2C%20%22opts%22%3A%20%22%20like%20%22%2C%20%22cols1%22%3A%20%22%22%7D%2C%20%225%22%3A%20%7B%22cols%22%3A%20%22atr%22%2C%20%22opts%22%3A%20%22%3E%22%2C%20%22strs%22%3A%20%2210%22%7D%2C%20%226%22%3A%20%7B%22cols%22%3A%20%22adx%22%2C%20%22opts%22%3A%20%22%3E%22%2C%20%22strs%22%3A%20%2225%22%7D%2C%20%227%22%3A%20%7B%22cols%22%3A%20%22avg_volume%22%2C%20%22opts%22%3A%20%22%3E%22%2C%20%22strs%22%3A%20%22500000%22%7D%2C%20%228%22%3A%20%7B%22cols%22%3A%20%22p_symbol%22%2C%20%22opts%22%3A%20%22%20not%20like%20%22%2C%20%22strs%22%3A%20%22%25-%25%22%7D%2C%20%229%22%3A%20%7B%22cols%22%3A%20%22%22%2C%20%22opts%22%3A%20%22%20like%20%22%2C%20%22strs%22%3A%20%22%22%7D%7D%7D&_=`)
	//fmt.Printf("printing data : %v", bytes.NewBuffer(requestBytes))
	// Create and modify HTTP request before sending
	request, err := http.NewRequest("POST", "https://www.icharts.in/includes/screener/EODScan.php", strings.NewReader(value))
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36")
	request.Header.Set("Accept-Language", "en-IN")
	request.Header.Set("Sec-Fetch-Site", "same-origin")
	request.Header.Set("Sec-Fetch-Mode", "cors")
	request.Header.Set("Sec-Fetch-Dest", "empty")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Content-Length", "77")
	request.Header.Set("X-Prototype-Version", "1.6.0.2")
	request.Header.Set("X-Requested-With", "XMLHttpRequest")
	request.Header.Set("Accept-Encoding", "br, gzip, deflate")
	request.Header.Set("Accept", "test/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	request.Header.Set("Referer", "https://www.icharts.in/screener-eod.html")
	request.Header.Set("Content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	request.Header.Set("Cookie", "__auc=de511acc174ab5077e9fdc45c58; __gads=ID=6614628d09480cf1:T=1600602042:S=ALNI_Ma7uQM537PiSjHCc-HJNkMjuCvB_g; __asc=8175300b174ac7e276e7f8e6040; __utma=192083122.261844769.1600602012.1600607119.1600621783.3; __utmc=192083122; __utmz=192083122.1600621783.3.2.utmcsr=google|utmccn=(organic)|utmcmd=organic|utmctr=(not%20provided); 0067889d30b7196cce14886855b1891a=10b2fd9745da54f7ed09c1d97cc59222; __utmb=192083122.2.10.1600621783; PHPSESSID=sc09fnbppc3qe2cfcerair9562")

	// Make request
	response, err := client.Do(request)
	//response, err := http.Get("https://www.icharts.in/screener-eod.html")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	data1, err := ioutil.ReadAll(response.Body)
	html_output := string(data1)
	fmt.Println("****************************************************")
	fmt.Printf("response : %+v \n", html_output)
	fmt.Println("****************************************************")

	htmlTokens := html.NewTokenizer(strings.NewReader(html_output))
	/*scrape := func(n *html.Node) {
		for _, attribute := range  n.Attr {
			fmt.Print(attribute.Key + " "+ attribute.Val + "\n")
			if "class" == attribute.Key && "mateTable" == attribute.Val {
				fmt.Println(n.Data)
			}
		}

	}
	fmt.Println("printing doc " + doc.Data)
	scrape(doc)*/
	fmt.Println("*******************BULLISH STOCKS*******************")
	fmt.Println("****************************************************")
	var stocksData [][]string
	var rowData []string
	isHeaderFound := false
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
			isRow  := t.Data == "tr"
			if isRow {
				//fmt.Printf("\n \n")
				if len(rowData) > 0 {

					stocksData = append(stocksData, rowData) //appending new row data to double array
					rowData = make([]string, 0)
				}else {
					fmt.Println("I am in else condition of len(rowData)")
					rowData = make([]string, 0)
				}


				_ = htmlTokens.Next() // navigate to next html token
				td := htmlTokens.Token()
				isAnchor := td.Data == "td" // check if current html tag is td

				if isAnchor { // if it is <td> tag then go ahead and extract data present inside <td>
					_ = htmlTokens.Next()
					eg := htmlTokens.Token()
					if isHeaderFound { // getting data from header that is <a> tag
						_ = htmlTokens.Next()
						header := htmlTokens.Token()
						log.Printf("data received in header inside row one: %s \n", header.Data)
						if header.Data == "a" {
							actualContent := header.String()
							rowData = append(rowData, actualContent) // appending extracted data to a row of data
							isHeaderFound = true
						}

					} else { // getting data from <td> tag
						//log.Printf(strconv.Itoa(index) + "." +strings.ReplaceAll(eg.String(), "&lt;\\/td&gt;", ""))
						actualString := utils.ReplaceUnnecessaryHtmlData(eg)
						if actualString != "" {
							rowData = append(rowData, actualString)
						}
					}
				}
				index++

				//for each column tag td present ger data present inside td tag and add it to the string slice
			} else if t.Data == "td" {
				_ = htmlTokens.Next()
				eg := htmlTokens.Token()
				if isHeaderFound { // getting data from header that is <a> tag
					_ = htmlTokens.Next()
					header := htmlTokens.Token()
					log.Printf("data received in header inside column: %s \n", header.Data)
					if header.Data == "a" {
						actualContent := header.String()
						rowData = append(rowData, actualContent)
						isHeaderFound = true
					}

				} else {
					//log.Printf(" | "+strings.ReplaceAll(eg.String(), "&lt;\\/td&gt;", ""))
					actualString := utils.ReplaceUnnecessaryHtmlData(eg)
					if actualString != "" {
						rowData = append(rowData, actualString)
					}
				}
			}
		}
	}

	csvFile, err := fileio.CreateCSVFile()
	if err != nil {
		log.Panic("error while creating csv file check for errors while creating")
	}
	err = fileio.WriteCSVFile(csvFile, stocksData)
	if err != nil {
		log.Panic("not able to write to csv file error while writing to csv file")
	}
}

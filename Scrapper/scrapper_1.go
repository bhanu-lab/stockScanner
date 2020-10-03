package Scrapper

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func ScrapeContent_2() {
	client := &http.Client{}
	var data = strings.NewReader(`json=%7B%22action%22%3A%20%22advanced_search%22%2C%20%22info%22%3A%20%7B%220%22%3A%20%7B%22cols%22%3A%20%22last_close%22%2C%20%22opts%22%3A%20%22%3E%22%2C%20%22cols1%22%3A%20%22ema_5%22%7D%2C%20%221%22%3A%20%7B%22cols%22%3A%20%22last_close%22%2C%20%22opts%22%3A%20%22%3E%22%2C%20%22cols1%22%3A%20%22ema_20%22%7D%2C%20%222%22%3A%20%7B%22cols%22%3A%20%22last_close%22%2C%20%22opts%22%3A%20%22%3C%22%2C%20%22cols1%22%3A%20%22bband_upper%22%7D%2C%20%223%22%3A%20%7B%22cols%22%3A%20%22%22%2C%20%22opts%22%3A%20%22%20like%20%22%2C%20%22cols1%22%3A%20%22%22%7D%2C%20%224%22%3A%20%7B%22cols%22%3A%20%22%22%2C%20%22opts%22%3A%20%22%20like%20%22%2C%20%22cols1%22%3A%20%22%22%7D%2C%20%225%22%3A%20%7B%22cols%22%3A%20%22atr%22%2C%20%22opts%22%3A%20%22%3E%22%2C%20%22strs%22%3A%20%2210%22%7D%2C%20%226%22%3A%20%7B%22cols%22%3A%20%22adx%22%2C%20%22opts%22%3A%20%22%3E%22%2C%20%22strs%22%3A%20%2225%22%7D%2C%20%227%22%3A%20%7B%22cols%22%3A%20%22avg_volume%22%2C%20%22opts%22%3A%20%22%3E%22%2C%20%22strs%22%3A%20%22500000%22%7D%2C%20%228%22%3A%20%7B%22cols%22%3A%20%22p_symbol%22%2C%20%22opts%22%3A%20%22%20not%20like%20%22%2C%20%22strs%22%3A%20%22%25-%25%22%7D%2C%20%229%22%3A%20%7B%22cols%22%3A%20%22%22%2C%20%22opts%22%3A%20%22%20like%20%22%2C%20%22strs%22%3A%20%22%22%7D%7D%7D&_=`)

	fmt.Printf("printing data : %v", data)

	req, err := http.NewRequest("POST", "https://www.icharts.in/includes/screener/EODScan.php", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept", "text/javascript, text/html, application/xml, text/xml, */*")
	req.Header.Set("X-Prototype-Version", "1.6.0.2")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36")
	req.Header.Set("Content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Origin", "https://www.icharts.in")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://www.icharts.in/screener-eod.html")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,te-IN;q=0.8,te;q=0.7,en-IN;q=0.6")
	req.Header.Set("Cookie", "__auc=de511acc174ab5077e9fdc45c58; __gads=ID=6614628d09480cf1:T=1600602042:S=ALNI_Ma7uQM537PiSjHCc-HJNkMjuCvB_g; __utmz=192083122.1600621783.3.2.utmcsr=google|utmccn=(organic)|utmcmd=organic|utmctr=(not%20provided); __utmc=192083122; PHPSESSID=8qim5nkje9947lgthl6ttggc35; 0067889d30b7196cce14886855b1891a=061d2caf23fb6a3e57eccf014ec54d52; __asc=5ab10303174b133e279579d4f63; __utma=192083122.261844769.1600602012.1600696007.1600700802.6; __utmt=1; __utmb=192083122.1.10.1600700802")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}

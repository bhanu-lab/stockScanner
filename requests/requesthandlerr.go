// requests provides making requests to other API's or http requests
package requests

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

/*
 * CreateAPIRequestAndGetResponse creates request with params encoded, converts response from request to string
 * returns error and string output of response
 */
func CreateAPIRequestAndGetResponse(value string, client *http.Client) (error, string) {
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
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	data1, err := ioutil.ReadAll(response.Body)
	htmlOutput := string(data1)
	return err, htmlOutput
}

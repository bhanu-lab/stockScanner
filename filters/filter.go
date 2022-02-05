// filters provides utilities for handling various utility operations
package filters

import (
	"log"
	"sort"
	"stockScanner/types"
	"strings"

	"golang.org/x/net/html"
)

/*
 *	ReplaceUnnecessaryHtmlData replaces unnecessary html data left over after extracting required data from html page
 *	uses strings.ReplaceAll() to replace old content with "" string	in given input parameter eg
 */
func ReplaceUnnecessaryHtmlData(eg html.Token) string {
	actualString := strings.ReplaceAll(eg.String(), "&lt;\\/td&gt;", "")
	actualString = strings.ReplaceAll(actualString, `<a href="\&#34;javascript:" toajaxtableeditor('order_by_changed',="" new="" array('`, "")
	actualString = strings.ReplaceAll(actualString, `','desc'));\"="">`, "")
	actualString = strings.ReplaceAll(actualString, `','asc'));\"="">`, "")
	return actualString
}

/* GetStockScannerType reads arguments when run and if option mentioned is -t and value is BEARISH/BULLISH
* accordingly types.BEARISH/ types.BULLISH will be returned if none is present or no option is selected then
* types.BULLISH will be selected by default
 */
func GetStockScannerType(args []string) int {
	//checks if given argument contains t then reads next stock option
	if strings.Contains(args[0], "t") {
		log.Printf("ouptut stock scanner type is selected")
		if strings.EqualFold(args[1], "Bearish") {
			return types.BEARISH
		}
	}
	return types.BULLISH
}

/*
 * GetCommandLineOptions takes all command line arguments as input
 * returns all the options chosen
 */
func GetCommandLineOptions(args []string) []string {
	var options []string

	for _, arg := range args {

		if strings.HasPrefix(arg, "-") {
			choice := strings.ReplaceAll(arg, "-", "")
			options = append(options, strings.Split(choice, "")...)
		}
	}

	log.Printf("options received are %+v \n", options)
	return options
}

/*
ConstructArgsMap - constructs a map with key as option and value as value from
command line arguments
*/
func ConstructArgsMap(args []string) map[string]string {
	argValues := make(map[string]string)

	for ind, arg := range args {

		if strings.HasPrefix(arg, "-") {
			choice := strings.ReplaceAll(arg, "-", "")
			if !(len(args) <= ind+1) && !strings.HasPrefix(args[ind+1], "-") {
				value := args[ind+1]
				argValues[choice] = value
			} else {
				argValues[choice] = ""
			}
		}
	}

	return argValues
}

// ValidateCommandLineOptions validates input parameters options provided
func ValidateCommandLineOptions(options []string) bool {
	if len(options) > 3 {
		log.Printf("options provided are more than allowed")
		return false
	}

	return true
}

/* FilterSimpleTechnicalValues filters output with only simple attribute values
 * INPUT : [][]string slice
 * OUTPUT : [][]string slice
 */
func FilterSimpleTechnicalValues(input [][]string) [][]string {
	output := make([][]string, 0)                          // create filtered output placeholder
	header := input[0]                                     // assign first row i.e header to compare against desired values
	columnsToBeIncluded := getIndexNumbersToFilter(header) // receives row index values which are to be included in filtered output
	sort.Ints(columnsToBeIncluded)
	for _, row := range input {
		filteredRow := make([]string, 0)
		for _, valIdx := range columnsToBeIncluded { // iterate over columns to be included from original data
			filteredRow = append(filteredRow, row[valIdx])
		}
		output = append(output, filteredRow)
	}
	return output
}

/*
 * getIndexNumbersToFilter returns all index numbers of fields matching simpleAttributes defined in this function
 * INPUT: slice string
 * OUTPUT: slice of intergers
 */
func getIndexNumbersToFilter(header []string) []int {
	var idxNums []int
	// predefined slice of header values which are to be included in simple output option
	simpleAttributes := []string{"p_symbol", "last_close", "avg_volume", "ema_8", "ema_20", "sma_50", "sma_200", "bband_upper", "bband_lower", "adx", "atr", "rsi"}
	for idx, headerValue := range header {
		if ok := IsValuePresentInStringSlice(headerValue, simpleAttributes); ok {
			idxNums = append(idxNums, idx) // prepare list of index numbers which are to be included
		}
	}
	return idxNums
}

/* IsValuePresentInStringSlice checks if string value present in given string slice
 * INPUT: value string, slice string
 * OUTPUT: boolean value
 */
func IsValuePresentInStringSlice(value string, list []string) bool {
	for _, element := range list {
		if ok := strings.EqualFold(value, element); ok {
			return true
		}
	}
	return false
}

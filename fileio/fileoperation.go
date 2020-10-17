// fileio provides different operations related to file operations
package fileio

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"stockScanner/types"
	"time"
)

/* CreateCSVFile creates csv file and returns file pointer with current time as file name in the format 20060102150405
 * GO follows below constants for formatting of date
	const (
    stdLongMonth      = "January"
    stdMonth          = "Jan"
    stdNumMonth       = "1"
    stdZeroMonth      = "01"
    stdLongWeekDay    = "Monday"
    stdWeekDay        = "Mon"
    stdDay            = "2"
    stdUnderDay       = "_2"
    stdZeroDay        = "02"
    stdHour           = "15"
    stdHour12         = "3"
    stdZeroHour12     = "03"
    stdMinute         = "4"
    stdZeroMinute     = "04"
    stdSecond         = "5"
    stdZeroSecond     = "05"
    stdLongYear       = "2006"
    stdYear           = "06"
    stdPM             = "PM"
    stdpm             = "pm"
    stdTZ             = "MST"
    stdISO8601TZ      = "Z0700"  // prints Z for UTC
    stdISO8601ColonTZ = "Z07:00" // prints Z for UTC
    stdNumTZ          = "-0700"  // always numeric
    stdNumShortTZ     = "-07"    // always numeric
    stdNumColonTZ     = "-07:00" // always numeric
)
* stockScannerType input parameter to read whether stock scanning option selected is bullish or bearish
*/
func CreateCSVFile(stockScannerType int) (*os.File, error) {
	log.Println("I am in create csv file")
	currentTime := time.Now()
	//timeValue := currentTime.Format("2017-08-21")
	timeValue := currentTime.Format("20060102150405")
	log.Printf("Current time is %s \n", timeValue)
	fileType := "Bullish_"
	if stockScannerType == types.BEARISH {
		fileType = "Bearish_"
	}
	path := "/home/bhanureddy/Documents/" + fileType + timeValue + ".csv"
	if err := os.MkdirAll(filepath.Dir(path), 0777); err != nil {
		log.Panic("Mentioned path is not available")
		return nil, err
	}

	fp, err := os.Create(path)
	if err != nil {
		log.Panic("error while creating file")
		return nil, err
	}
	return fp, nil
}

/*
* WriteCSVFile writes to csv file with given string slice as input
* return error if there is any while opening file or writing data to file
* return nil if write operation is successful
 */
func WriteCSVFile(fp *os.File, data [][]string) error {
	/*log.Printf("Received data to write to csv is %#v \n", data)*/
	log.Println("Wrote data to ", fp.Name())
	file, err := os.OpenFile(fp.Name(), os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Panic("not able to open file :" + file.Name())
		return err
	}
	csvWriter := csv.NewWriter(file)
	err = csvWriter.WriteAll(data)
	if err != nil {
		log.Panic("Not able to write data to file : " + file.Name())
		return err
	}
	fmt.Print("****************************DONE*****************************")
	return nil
}

package fileio

import (
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
	"time"
)

// CreateCSVFile creates csv file and returns file pointer with current time as file name in this format "2017_08_21_17_06_19"
func CreateCSVFile() (*os.File, error){
	log.Println("I am in create csv file")
	currentTime := time.Now()
	//timeValue := currentTime.Format("2017-08-21")
	timeValue := currentTime.Format("20060102150405")
	log.Printf("Current time is %s \n", timeValue)
	path := "/home/bhanureddy/Documents/"+timeValue+".csv"
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

// WriteCSVFile writes to csv file with given string slice as input
// return error if there is any while opening file or writing data to file
// return nil if write operation is successfull
func WriteCSVFile(fp *os.File, data [][]string) error {
	log.Printf("Received data to write to csv is %#v \n", data)
	file, err := os.OpenFile(fp.Name(), os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Panic("not able to open file :" + file.Name())
		return err
	}
	csvWriter := csv.NewWriter(file)
	err = csvWriter.WriteAll(data)
	if err != nil {
		log.Panic("Not able to write data to file : "+file.Name())
		return err
	}
	return nil
}

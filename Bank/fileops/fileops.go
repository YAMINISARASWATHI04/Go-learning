package fileops


import (
	"errors"
	"fmt"
	"os"
	"strconv"
	
)



func WriteFloatToFile(value float64,filename string) {
	valueText := fmt.Sprint(value)
	os.WriteFile("balance.txt", []byte(valueText), 0644) // converts to byte collection
	// number stands for File permissions notation
	// Owners can read and write where as others have only read permissions
}


func GetFloatfromfile(filename string) (float64, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return 1000, errors.New("Failed to read file")
	} // If there is an issue in reading the file
	// nil stands for absence of a useful value
	valuetext := string(data)
	val, err := strconv.ParseFloat(valuetext, 64)
	if err != nil {
		return 1000, errors.New("Failed to convert into specified type")
	} // If there is an issue in converting to number
	return val, nil
}
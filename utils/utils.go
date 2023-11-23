package utils

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func GetRaw(source map[string]int, bitarray []bool) int {
	offset := source["offset"]
	size := source["size"]

	slicedBits := bitarray[offset : offset+size]

	var binaryString string
	for _, digit := range slicedBits {
		if digit {
			binaryString += "1"
		} else {
			binaryString += "0"
		}
	}

	rawData, err := strconv.ParseInt(binaryString, 2, 64)
	if err != nil {
		fmt.Println("Error converting binary to integer:", err)
		return 0 // or handle the error in an appropriate way
	}

	return int(rawData)
}

func CombineHex(data []int) int {
	output := 0x00
	for i, value := range reverseIntSlice(data) {
		output |= (value << (i * 8))
	}
	return output
}

func ToBitArray(data interface{}, width int) []bool {
	var combinedData int
	switch v := data.(type) {
	case []int:
		combinedData = CombineHex(v)
	case []byte:
		combinedData = CombineHex(byteArrayToIntSlice(v))
	case int:
		combinedData = v
	default:
		fmt.Println("Unsupported data type")
		return nil
	}

	binaryString := fmt.Sprintf("%b", combinedData)
	paddedBinaryString := fmt.Sprintf("%0*s", width, binaryString)

	var bitArray []bool
	for _, digit := range paddedBinaryString {
		bitArray = append(bitArray, digit == '1')
	}
	return bitArray
}

func reverseIntSlice(slice []int) []int {
	length := len(slice)
	reversed := make([]int, length)
	for i, value := range slice {
		reversed[length-i-1] = value
	}
	return reversed
}

func byteArrayToIntSlice(byteArray []byte) []int {
	intSlice := make([]int, len(byteArray))
	for i, b := range byteArray {
		intSlice[i] = int(b)
	}
	return intSlice
}

func LoadXML(pathToXml string) []byte {
	xmlFile, err := os.Open(pathToXml)
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()
	loadedProfiles, _ := io.ReadAll(xmlFile)
	if err != nil {
		fmt.Println(err)
	}
	return loadedProfiles
}

func LoadJSON(pathToJSON string) []byte {
	JSONFile, err := os.Open(pathToJSON)
	if err != nil {
		fmt.Println(err)
	}
	defer JSONFile.Close()
	loadedProfiles, _ := io.ReadAll(JSONFile)
	if err != nil {
		fmt.Println(err)
	}
	return loadedProfiles
}

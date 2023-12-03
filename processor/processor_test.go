package processor

import (
	"testing"

	"github.com/anesmemisevic/enocean-transcoder/utils"
)

func TestBytesToBitArrayD21441(t *testing.T) {
	byteArray := []byte{139, 78, 197, 57, 5, 121, 194, 125, 17}
	bitArray := utils.ToBitArray(byteArray)
	expectedBitArray := []bool{true, false, false, false, true, false, true, true, false, true, false, false, true, true, true, false, true, true, false, false, false, true, false, true, false, false, true, true, true, false, false, true, false, false, false, false, false, true, false, true, false, true, true, true, true, false, false, true, true, true, false, false, false, false, true, false, false, true, true, true, true, true, false, true, false, false, false, true, false, false, false, true}
	if len(bitArray) != len(expectedBitArray) {
		t.Errorf("BitArray length is not equal to expectedBitArray length")
	}
	for i := 0; i < len(bitArray); i++ {
		if bitArray[i] != expectedBitArray[i] {
			t.Errorf("BitArray is not equal to expectedBitArray")
		}
	}
}

func TestBytesToBitArrayD21441Negative(t *testing.T) {
	byteArray := []byte{139, 78, 197, 57, 5, 121, 194, 125, 18}
	bitArray := utils.ToBitArray(byteArray)
	expectedBitArray := []bool{true, false, false, false, true, false, true, true, false, true, false, false, true, true, true, false, true, true, false, false, false, true, false, true, false, false, true, true, true, false, false, true, false, false, false, false, false, true, false, true, false, true, true, true, true, false, false, true, true, true, false, false, false, false, true, false, false, true, true, true, true, true, false, true, false, false, false, true, false, false, false, true}

	if len(bitArray) != len(expectedBitArray) {
		t.Errorf("BitArray length is not equal to expectedBitArray length")
	}

	// negative test case: should fail
	incorrectValue := 0
	for i := 0; i < len(bitArray); i++ {
		if bitArray[i] == expectedBitArray[i] {
			incorrectValue++
		}
	}

	if incorrectValue > 0 {
		t.Logf("BitArray is not equal to expectedBitArray, %d values are incorrect", incorrectValue)
	}
}

func TestIntsToBitArrayD21441(t *testing.T) {
	byteArray := []int{139, 78, 197, 57, 5, 121, 194, 125, 17}
	bitArray := utils.ToBitArray(byteArray)
	expectedBitArray := []bool{true, false, false, false, true, false, true, true, false, true, false, false, true, true, true, false, true, true, false, false, false, true, false, true, false, false, true, true, true, false, false, true, false, false, false, false, false, true, false, true, false, true, true, true, true, false, false, true, true, true, false, false, false, false, true, false, false, true, true, true, true, true, false, true, false, false, false, true, false, false, false, true}
	if len(bitArray) != len(expectedBitArray) {
		t.Errorf("BitArray length is not equal to expectedBitArray length")
	}
	for i := 0; i < len(bitArray); i++ {
		if bitArray[i] != expectedBitArray[i] {
			t.Errorf("BitArray is not equal to expectedBitArray")
		}
	}
}

func TestLoadEEPs(t *testing.T) {
	eeps := LoadEEPs()

	if len(eeps.Telegrams) == 0 {
		t.Errorf("EEPs not loaded")
	}

	if len(eeps.Telegrams[0].Profiles) == 0 {
		t.Errorf("EEPs not loaded")
	}

	if len(eeps.Telegrams[0].Profiles[0].Profiles) == 0 {
		t.Errorf("EEPs not loaded")
	}

	if len(eeps.Telegrams[0].Profiles[0].Profiles[0].Data) == 0 {
		t.Errorf("EEPs not loaded")
	}

}

func TestFindProfileD21441(t *testing.T) {

	findRorg := "0xD2"
	findFunc := "0x14"
	findType := "0x41"

	eeps := LoadEEPs()
	profile, ok := FindProfile(eeps, findRorg, findFunc, findType)

	if !ok {
		t.Errorf("Profile not found")
	}

	if profile.Type != findType {
		t.Errorf("Profile with wrong type found")
	}

	if profile.Description != "Indoor -Temperature, Humidity XYZ Acceleration, Illumination Sensor" {
		t.Errorf("Profile not found")
	}

}

func TestLoadSensorValuesMetadataHumTempValue(t *testing.T) {
	findRorg := "0xD2"
	findFunc := "0x14"
	findType := "0x41"

	eeps := LoadEEPs()
	profile, ok := FindProfile(eeps, findRorg, findFunc, findType)

	if !ok {
		t.Errorf("Profile not found")
	}

	dataMap := LoadSensorValuesMetadata(profile)

	if len(dataMap) == 0 {
		t.Errorf("Data map not loaded")
	}

	if dataMap["TMP"].(map[string]interface{})["datatype"] != "value" {
		t.Errorf("Datatype not correct")
	}

	if dataMap["TMP"].(map[string]interface{})["description"] != "Temperature 10" {
		t.Errorf("Description not correct")
	}

	if dataMap["TMP"].(map[string]interface{})["shortcut"] != "TMP" {
		t.Errorf("Shortcut not correct")
	}

	if dataMap["TMP"].(map[string]interface{})["offset"] != 0 {
		t.Errorf("Offset not correct")
	}

	if dataMap["TMP"].(map[string]interface{})["size"] != 10 {
		t.Errorf("Size not correct")
	}

	if dataMap["TMP"].(map[string]interface{})["unit"] != "°C" {
		t.Errorf("Unit not correct")
	}

	if dataMap["TMP"].(map[string]interface{})["range"].(map[string]interface{})["max"] != float64(1000) {
		t.Errorf("Temperature range max not correct")
	}

	if dataMap["TMP"].(map[string]interface{})["range"].(map[string]interface{})["min"] != float64(0) {
		t.Errorf("Temperature range min not correct")
	}

	if dataMap["TMP"].(map[string]interface{})["scale"].(map[string]interface{})["max"] != 60.0 {
		t.Errorf("Temperature scale max not correct")
	}

	if dataMap["TMP"].(map[string]interface{})["scale"].(map[string]interface{})["min"] != -40.0 {
		t.Errorf("Temperature scale min not correct")
	}

	if dataMap["HUM"].(map[string]interface{})["description"] != "Rel. Humidity linear)" {
		t.Errorf("Humidity description not correct")
	}

	if dataMap["HUM"].(map[string]interface{})["shortcut"] != "HUM" {
		t.Errorf("Humidity shortcut not correct")
	}

	if dataMap["HUM"].(map[string]interface{})["offset"] != 10 {
		t.Errorf("Humidity offset not correct")
	}

	if dataMap["HUM"].(map[string]interface{})["size"] != 8 {
		t.Errorf("Humidity size not correct")
	}

	if dataMap["HUM"].(map[string]interface{})["unit"] != "%" {
		t.Errorf("Humidity unit not correct")
	}

	if dataMap["HUM"].(map[string]interface{})["range"].(map[string]interface{})["max"] != 200.0 {
		t.Errorf("Humidity range max not correct")
	}

	if dataMap["HUM"].(map[string]interface{})["range"].(map[string]interface{})["min"] != 0.0 {
		t.Errorf("Humidity range min not correct")
	}

	if dataMap["HUM"].(map[string]interface{})["scale"].(map[string]interface{})["max"] != 100.0 {
		t.Errorf("Humidity scale max not correct")
	}

	if dataMap["HUM"].(map[string]interface{})["scale"].(map[string]interface{})["min"] != 0.0 {
		t.Errorf("Humidity scale min not correct")
	}

}

func TestLoadSensorValuesMetadataForContactEnum(t *testing.T) {
	findRorg := "0xD2"
	findFunc := "0x14"
	findType := "0x41"

	eeps := LoadEEPs()
	profile, ok := FindProfile(eeps, findRorg, findFunc, findType)

	if !ok {
		t.Errorf("Profile not found")
	}

	dataMap := LoadSensorValuesMetadata(profile)

	if len(dataMap) == 0 {
		t.Errorf("Data map not loaded")
	}

	if dataMap["CO"].(map[string]interface{})["datatype"] != "enum" {
		t.Errorf("Contact datatype not correct")
	}

	if dataMap["CO"].(map[string]interface{})["description"] != "Contact" {
		t.Errorf("Contact description not correct")
	}

	if dataMap["CO"].(map[string]interface{})["shortcut"] != "CO" {
		t.Errorf("Contact shortcut not correct")
	}

	if dataMap["CO"].(map[string]interface{})["offset"] != 67 {
		t.Errorf("Contact offset not correct")
	}

	if dataMap["CO"].(map[string]interface{})["size"] != 1 {
		t.Errorf("Contact size not correct")
	}

	if (dataMap["CO"].(map[string]interface{})["item"].(map[string]interface{})["Closed"]) != 1 {
		t.Errorf("Contact item not correct")
	}

	if (dataMap["CO"].(map[string]interface{})["item"].(map[string]interface{})["Open"]) != 0 {
		t.Errorf("Contact item not correct")
	}
}

func TestLoadSensorValuesMetadataForIlluminationValue(t *testing.T) {
	findRorg := "0xD2"
	findFunc := "0x14"
	findType := "0x41"

	eeps := LoadEEPs()
	profile, ok := FindProfile(eeps, findRorg, findFunc, findType)

	if !ok {
		t.Errorf("Profile not found")
	}

	dataMap := LoadSensorValuesMetadata(profile)

	if len(dataMap) == 0 {
		t.Errorf("Data map not loaded")
	}

	if dataMap["ILL"].(map[string]interface{})["description"] != "Illumination linear)" {
		t.Errorf("Illumination description not correct")
	}

	if dataMap["ILL"].(map[string]interface{})["shortcut"] != "ILL" {
		t.Errorf("Illumination shortcut not correct")
	}

	if dataMap["ILL"].(map[string]interface{})["offset"] != 18 {
		t.Errorf("Illimination offset not correct")
	}

	if dataMap["ILL"].(map[string]interface{})["size"] != 17 {
		t.Errorf("Illimination size not correct")
	}

	if dataMap["ILL"].(map[string]interface{})["unit"] != "lx" {
		t.Errorf("Illimination unit not correct")
	}

	if dataMap["ILL"].(map[string]interface{})["range"].(map[string]interface{})["max"] != 100000.0 {
		t.Errorf("Illimination range max not correct")
	}

	if dataMap["ILL"].(map[string]interface{})["range"].(map[string]interface{})["min"] != 0.0 {
		t.Errorf("Illimination range min not correct")
	}

	if dataMap["ILL"].(map[string]interface{})["scale"].(map[string]interface{})["max"] != 100000.0 {
		t.Errorf("Illimination scale max not correct")
	}

	if dataMap["ILL"].(map[string]interface{})["scale"].(map[string]interface{})["min"] != 0.0 {
		t.Errorf("Illimination scale min not correct")
	}
}

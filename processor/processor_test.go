package processor

import (
	"fmt"
	"testing"

	"github.com/anesmemisevic/enocean-transcoder/utils"
)

func TestToBitArrayD21441(t *testing.T) {
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

func TestLoadEEPs(t *testing.T) {
	findRorg := "0xD2"
	findFunc := "0x14"
	findType := "0x41"

	eeps := LoadEEPs(findRorg, findFunc, findType)

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

	eeps := LoadEEPs(findRorg, findFunc, findType)
	ok, profile := FindProfile(eeps, findRorg, findFunc, findType)

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

func TestLoadSensorValuesMetadata(t *testing.T) {
	findRorg := "0xD2"
	findFunc := "0x14"
	findType := "0x41"

	eeps := LoadEEPs(findRorg, findFunc, findType)
	ok, profile := FindProfile(eeps, findRorg, findFunc, findType)

	if !ok {
		t.Errorf("Profile not found")
	}

	dataMap := LoadSensorValuesMetadata(profile)

	if len(dataMap) == 0 {
		t.Errorf("Data map not loaded")
	}

	if dataMap["TMP"].(map[string]interface{})["datatype"] != "value" {
		t.Errorf("Data map not loaded")
	}

	if dataMap["TMP"].(map[string]interface{})["description"] != "Temperature 10" {
		t.Errorf("Data map not loaded")
	}

	if dataMap["TMP"].(map[string]interface{})["shortcut"] != "TMP" {
		t.Errorf("Data map not loaded")
	}

	if dataMap["TMP"].(map[string]interface{})["offset"] != 0 {
		t.Errorf("Data map not loaded")
	}

	if dataMap["TMP"].(map[string]interface{})["size"] != 10 {
		t.Errorf("Data map not loaded")
	}

	if dataMap["TMP"].(map[string]interface{})["unit"] != "°C" {
		t.Errorf("Data map not loaded")
	}

	if dataMap["TMP"].(map[string]interface{})["range"].(map[string]interface{})["max"] != float64(1000) {
		fmt.Println(dataMap["TMP"].(map[string]interface{})["range"].(map[string]interface{})["max"])
		t.Errorf("Data map not loaded")
	}

	if dataMap["TMP"].(map[string]interface{})["range"].(map[string]interface{})["min"] != float64(0) {
		t.Errorf("Data map not loaded")
	}

	if dataMap["TMP"].(map[string]interface{})["scale"].(map[string]interface{})["max"] != 60.0 {
		t.Errorf("Data map not loaded")
	}

	if dataMap["TMP"].(map[string]interface{})["scale"].(map[string]interface{})["min"] != -40.0 {
		t.Errorf("Data map not loaded")
	}
}

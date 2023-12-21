package transcoder

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/anesmemisevic/enocean-transcoder/models"
	"github.com/anesmemisevic/enocean-transcoder/processor"
	"github.com/anesmemisevic/enocean-transcoder/utils"
)

type SensorDecoder struct {
	ByteStream  []int
	EEP         models.EEP
	EEPMetadata models.Telegrams
}

func (f *SensorDecoder) LoadEEPs() (loaded bool) {
	if len(f.EEPMetadata.Telegrams) > 0 {
		return true
	}
	loadedProfiles := utils.LoadXML("EEP.xml")
	var telegrams models.Telegrams
	ok := xml.Unmarshal(loadedProfiles, &telegrams)
	if ok != nil {
		fmt.Println("Error loading EEP.xml")
	}
	f.EEPMetadata = telegrams
	return len(f.EEPMetadata.Telegrams) > 0
}

func (f SensorDecoder) Decode(byteStream []int, eep models.EEP) (sensor models.Sensor, ok bool) {
	f.ByteStream = byteStream
	bitArray := utils.ToBitArray(f.ByteStream)
	profile, ok := processor.FindProfile(f.EEPMetadata, eep.Rorg, eep.Func, eep.Type)
	if !ok {
		fmt.Println("Error processing telegram")
	}
	eepSensorMetadata := processor.LoadSensorValuesMetadata(profile)
	sensorValues, ok := processor.GetSensorValues(eepSensorMetadata, bitArray)
	if !ok {
		fmt.Println("Error processing telegram")
	}

	sensorData := make(map[string]models.SensorData)
	for key, value := range sensorValues {
		sensorData[key] = models.SensorData{
			Shortcut:    key,
			Value:       value.(map[string]interface{})["scaledValue"],
			Unit:        value.(map[string]interface{})["unit"],
			Description: value.(map[string]interface{})["description"].(string),
		}
	}
	sensor = models.Sensor{
		AssignedEEP:       eep,
		Data:              sensorData,
		TelegramType:      f.EEPMetadata.Telegrams[0].Type,
		SensorDescription: profile.Description,
	}

	return sensor, true

}

func MarshalJSON(sensor models.Sensor) (jsonData []byte, err error) {
	type Alias models.Sensor
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(&sensor),
	})
}

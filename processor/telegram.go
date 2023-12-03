package processor

import (
	"encoding/xml"
	"fmt"

	"github.com/anesmemisevic/enocean-transcoder/models"
	"github.com/anesmemisevic/enocean-transcoder/utils"
)

func LoadEEPs() models.Telegrams {
	loadedProfiles := utils.LoadXML("EEP.xml")
	var telegrams models.Telegrams
	ok := xml.Unmarshal(loadedProfiles, &telegrams)
	if ok != nil {
		fmt.Println("Error loading EEP.xml")
	}
	return telegrams
}

func FindProfile(telegrams models.Telegrams, targetRorg string, targetFunc string, targetType string) (result models.Profile, found bool) {
	for _, telegram := range telegrams.Telegrams {
		if telegram.Rorg == targetRorg {
			for _, profile := range telegram.Profiles {
				if profile.Func == targetFunc {
					for _, profileData := range profile.Profiles {
						if profileData.Type == targetType {
							return profileData, true
						}
					}
				}
			}
		}
	}
	return models.Profile{}, false
}

func LoadSensorValuesMetadata(profile models.Profile) (SensorValuesMetadataMap map[string]interface{}) {
	valuesMetadataMap := make(map[string]interface{})

	for _, data := range profile.Data {
		if len(data.Value) != 0 {
			for _, value := range data.Value {
				valuesMetadataMap[value.Shortcut] = make(map[string]interface{})
				valuesMetadataMap[value.Shortcut].(map[string]interface{})["datatype"] = "value"
				valuesMetadataMap[value.Shortcut].(map[string]interface{})["description"] = value.Description
				valuesMetadataMap[value.Shortcut].(map[string]interface{})["shortcut"] = value.Shortcut
				valuesMetadataMap[value.Shortcut].(map[string]interface{})["offset"] = value.Offset
				valuesMetadataMap[value.Shortcut].(map[string]interface{})["size"] = value.Size
				valuesMetadataMap[value.Shortcut].(map[string]interface{})["unit"] = value.Unit
				valuesMetadataMap[value.Shortcut].(map[string]interface{})["range"] = make(map[string]interface{})
				valuesMetadataMap[value.Shortcut].(map[string]interface{})["scale"] = make(map[string]interface{})
				valuesMetadataMap[value.Shortcut].(map[string]interface{})["range"].(map[string]interface{})["max"] = value.Range.Max
				valuesMetadataMap[value.Shortcut].(map[string]interface{})["range"].(map[string]interface{})["min"] = value.Range.Min
				valuesMetadataMap[value.Shortcut].(map[string]interface{})["scale"].(map[string]interface{})["max"] = value.Scale.Max
				valuesMetadataMap[value.Shortcut].(map[string]interface{})["scale"].(map[string]interface{})["min"] = value.Scale.Min
			}
		}

		if len(data.Status) != 0 {
			for _, status := range data.Status {
				valuesMetadataMap[status.Shortcut] = make(map[string]interface{})
				valuesMetadataMap[status.Shortcut].(map[string]interface{})["datatype"] = "status"
				valuesMetadataMap[status.Shortcut].(map[string]interface{})["description"] = status.Description
				valuesMetadataMap[status.Shortcut].(map[string]interface{})["shortcut"] = status.Shortcut
				valuesMetadataMap[status.Shortcut].(map[string]interface{})["offset"] = status.Offset
				valuesMetadataMap[status.Shortcut].(map[string]interface{})["size"] = status.Size
			}
		}

		if len(data.Enum) != 0 {
			for _, enum := range data.Enum {
				valuesMetadataMap[enum.Shortcut] = make(map[string]interface{})
				valuesMetadataMap[enum.Shortcut].(map[string]interface{})["datatype"] = "enum"
				valuesMetadataMap[enum.Shortcut].(map[string]interface{})["description"] = enum.Description
				valuesMetadataMap[enum.Shortcut].(map[string]interface{})["shortcut"] = enum.Shortcut
				valuesMetadataMap[enum.Shortcut].(map[string]interface{})["offset"] = enum.Offset
				valuesMetadataMap[enum.Shortcut].(map[string]interface{})["size"] = enum.Size
				valuesMetadataMap[enum.Shortcut].(map[string]interface{})["item"] = make(map[string]interface{})
				valuesMetadataMap[enum.Shortcut].(map[string]interface{})["rangeitem"] = make(map[string]interface{})

				if len(enum.Item) != 0 {
					for _, item := range enum.Item {
						valuesMetadataMap[enum.Shortcut].(map[string]interface{})["item"].(map[string]interface{})[item.Description] = item.Value
					}
				}

				if len(enum.RangeItem) != 0 {
					for _, rangeItem := range enum.RangeItem {
						valuesMetadataMap[enum.Shortcut].(map[string]interface{})["rangeitem"].(map[string]interface{})[rangeItem.Description+"-start"] = rangeItem.Start
						valuesMetadataMap[enum.Shortcut].(map[string]interface{})["rangeitem"].(map[string]interface{})[rangeItem.Description+"-end"] = rangeItem.End
					}
				}
			}
		}
	}
	return valuesMetadataMap
}

func GetSensorValues(dataMap map[string]interface{}, bitArray []bool) (sensorValues map[string]interface{}, ok bool) {
	sensorValues = make(map[string]interface{})
	for key, value := range dataMap {
		if value.(map[string]interface{})["datatype"] == "value" {
			// TODO: make new struct for sensor values including rawValue, description, shortcut, unit, realValue, etc.
			offset := value.(map[string]interface{})["offset"].(int)
			size := value.(map[string]interface{})["size"].(int)
			minScale := value.(map[string]interface{})["scale"].(map[string]interface{})["min"].(float64)
			maxScale := value.(map[string]interface{})["scale"].(map[string]interface{})["max"].(float64)
			minRange := value.(map[string]interface{})["range"].(map[string]interface{})["min"].(float64)
			maxRange := value.(map[string]interface{})["range"].(map[string]interface{})["max"].(float64)
			offsetSizeMap := map[string]int{"offset": offset, "size": size}

			rawValue := utils.GetRaw(offsetSizeMap, bitArray)
			sensorValues[key] = make(map[string]interface{})
			sensorValues[key].(map[string]interface{})["rawValue"] = rawValue
			sensorValues[key].(map[string]interface{})["description"] = value.(map[string]interface{})["description"]
			sensorValues[key].(map[string]interface{})["unit"] = value.(map[string]interface{})["unit"]
			sensorValues[key].(map[string]interface{})["scaledValue"] = utils.GetScaledValue(minScale, maxScale, minRange, maxRange, rawValue)
		}

		if value.(map[string]interface{})["datatype"] == "status" {
			offset := value.(map[string]interface{})["offset"].(int)
			size := value.(map[string]interface{})["size"].(int)
			offsetSizeMap := map[string]int{"offset": offset, "size": size}

			rawValue := utils.GetRaw(offsetSizeMap, bitArray)
			sensorValues[key] = make(map[string]interface{})
			sensorValues[key].(map[string]interface{})["description"] = value.(map[string]interface{})["description"]
			if rawValue == 1 {
				sensorValues[key].(map[string]interface{})["scaledValue"] = true
			} else {
				sensorValues[key].(map[string]interface{})["scaledValue"] = false
			}
		}

		if value.(map[string]interface{})["datatype"] == "enum" {
			offset := value.(map[string]interface{})["offset"].(int)
			size := value.(map[string]interface{})["size"].(int)
			offsetSizeMap := map[string]int{"offset": offset, "size": size}

			rawValue := utils.GetRaw(offsetSizeMap, bitArray)
			sensorValues[key] = make(map[string]interface{})
			sensorValues[key].(map[string]interface{})["unit"] = value.(map[string]interface{})["unit"]
			sensorValues[key].(map[string]interface{})["description"] = value.(map[string]interface{})["description"]
			sensorValues[key].(map[string]interface{})["rawValue"] = rawValue

			if len(value.(map[string]interface{})["item"].(map[string]interface{})) != 0 {
				for ItemKey, ItemValue := range value.(map[string]interface{})["item"].(map[string]interface{}) {
					if ItemValue == rawValue {
						sensorValues[key].(map[string]interface{})["scaledValue"] = ItemKey
					}
				}
			}

			if len(value.(map[string]interface{})["rangeitem"].(map[string]interface{})) != 0 {
				for RangeItemKey, RangeItemValue := range value.(map[string]interface{})["rangeitem"].(map[string]interface{}) {
					if rawValue >= RangeItemValue.(map[string]interface{})[RangeItemKey+"-start"].(int) && rawValue <= RangeItemValue.(map[string]interface{})[RangeItemKey+"-end"].(int) {
						sensorValues[key].(map[string]interface{})["scaledValue"] = RangeItemKey
					}
				}
			}
		}
	}

	if len(sensorValues) != 0 {
		return sensorValues, true
	}

	return nil, false
}

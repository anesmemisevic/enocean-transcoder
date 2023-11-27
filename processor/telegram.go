package processor

import (
	"encoding/xml"

	"github.com/anesmemisevic/enocean-transcoder/models"
	"github.com/anesmemisevic/enocean-transcoder/utils"
)

func LoadEEPs(rorg string, eepFunc string, eepType string) models.Telegrams {
	loadedProfiles := utils.LoadXML("EEP.xml")
	var telegrams models.Telegrams
	xml.Unmarshal(loadedProfiles, &telegrams)
	return telegrams
}

func FindProfile(telegrams models.Telegrams, targetRorg string, targetFunc string, targetType string) (found bool, result models.Profile) {
	for _, telegram := range telegrams.Telegrams {
		if telegram.Rorg == targetRorg {
			for _, profile := range telegram.Profiles {
				if profile.Func == targetFunc {
					for _, profileData := range profile.Profiles {
						if profileData.Type == targetType {
							return true, profileData
						}
					}
				}
			}
		}
	}
	return false, models.Profile{}
}

func LoadSensorValuesMetadata(profile models.Profile) (SensorValuesMetadataMap map[string]interface{}) {
	valuesMetadataMap := make(map[string]interface{})

	for _, data := range profile.Data {
		if len(data.Value) != 0 {
			for _, value := range data.Value {
				valuesMetadataMap[value.Shortcut] = make(map[string]interface{})
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
				valuesMetadataMap[status.Shortcut].(map[string]interface{})["description"] = status.Description
				valuesMetadataMap[status.Shortcut].(map[string]interface{})["shortcut"] = status.Shortcut
				valuesMetadataMap[status.Shortcut].(map[string]interface{})["offset"] = status.Offset
				valuesMetadataMap[status.Shortcut].(map[string]interface{})["size"] = status.Size
			}
		}

		if len(data.Enum) != 0 {
			for _, enum := range data.Enum {
				valuesMetadataMap[enum.Shortcut] = make(map[string]interface{})
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

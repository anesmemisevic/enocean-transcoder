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

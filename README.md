# EnOcean Serial Protocol 3 Transcoder for Go

![Go version](https://img.shields.io/badge/Go-1.21.4-blue.svg)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

## Overview

This project provides a Go implementation of a transcoder for the Data part of the EnOcean Serial Protocol 3 (ESP3) package. The transcoder allows you to decode the byte stream of data into human-readable information, making it easier to work with EnOcean devices in your Go applications.

## Features

- **Go Implementation:** Developed using Go, ensuring efficiency, simplicity, and ease of integration into your Go projects.
- **Decoding ESP3 Data:** Convert the byte stream from EnOcean devices into a structured, human-readable format.
- **Encoding ESP3 Data (In Implementation):** Convert structured, human-readable data into the byte stream format of the EnOcean Serial Protocol 3.

## Getting Started

**Note: This section is under construction, and specific details, including how to import the package into your project, will be added shortly. Please check back for updates.**

1. **Install Go:**
   Make sure you have Go version 1.20.11 or later installed on your system.

2. **Clone the Repository:**
   ```bash
   git clone https://github.com/anesmemisevic/enocean-transcoder.git

3. **Import in your project:**
  TBA
4. **Decode EnOcean Data:**
  TBA
5. **Encode EnOcean Data:**
  TBA


#### Current Version Example

```Go
package main

import (
	"fmt"

	"github.com/anesmemisevic/enocean-transcoder/processor"
	"github.com/anesmemisevic/enocean-transcoder/utils"
)

func main() {

	byteArrayMultisensor := []byte{139, 78, 197, 57, 5, 121, 194, 125, 17}
	bitArray := utils.ToBitArray(byteArrayMultisensor)
	findRorg := "0xD2"
	findFunc := "0x14"
	findType := "0x41"

	eeps := processor.LoadEEPs()
	profile, ok := processor.FindProfile(eeps, findRorg, findFunc, findType)
	if !ok {
		fmt.Println("Error processing telegram")
	}
	dataMap := processor.LoadSensorValuesMetadata(profile)
	valuesMap, ok := processor.GetSensorValues(dataMap, bitArray)
	if !ok {
		fmt.Println("Error processing telegram")
	}
	fmt.Println(valuesMap)
}
```

```json
{
  "assigned_eep": { "rorg": "0xD2", "func": "0x14", "type": "0x41" },
  "telegram_type": "VLD",
  "sensor_description": "Indoor -Temperature, Humidity XYZ Acceleration, Illumination Sensor",
  "data": {
    "ACC": {
      "shortcut": "ACC",
      "value": "Periodic Update",
      "unit": null,
      "description": "Acceleration Status"
    },
    "ACX": {
      "shortcut": "ACX",
      "value": 1,
      "unit": "g",
      "description": "Absolute Acceleration on X axis"
    },
    "ACY": {
      "shortcut": "ACY",
      "value": 2,
      "unit": "g",
      "description": "Absolute Acceleration on Y axis"
    },
    "ACZ": {
      "shortcut": "ACZ",
      "value": 2.5,
      "unit": "g",
      "description": "Absolute Acceleration on Z axis"
    },
    "CO": {
      "shortcut": "CO",
      "value": "Closed",
      "unit": null,
      "description": "Contact"
    },
    "HUM": {
      "shortcut": "HUM",
      "value": 29.5,
      "unit": "%",
      "description": "Rel. Humidity linear)"
    },
    "ILL": {
      "shortcut": "ILL",
      "value": 10696,
      "unit": "lx",
      "description": "Illumination linear)"
    },
    "TMP": {
      "shortcut": "TMP",
      "value": -31.900000000000006,
      "unit": "°C",
      "description": "Temperature 10"
    }
  }
}

```

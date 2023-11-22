package models

import (
	"encoding/xml"
)

type Telegrams struct {
	XMLName      xml.Name   `xml:"telegrams"`
	Telegrams    []Telegram `xml:"telegram"`
	Version      string     `xml:"version,attr,omitempty"`
	MajorVersion string     `xml:"major_version,attr,omitempty"`
	MinorVersion string     `xml:"minor_version,attr,omitempty"`
	Revision     string     `xml:"revision,attr,omitempty"`
}
type Telegram struct {
	XMLName     xml.Name   `xml:"telegram"`
	Profiles    []Profiles `xml:"profiles,omitempty"`
	Rorg        string     `xml:"rorg,attr,omitempty"`
	Type        string     `xml:"type,attr,omitempty"`
	Description string     `xml:"description,attr,omitempty"`
}

type Profiles struct {
	XMLName     xml.Name  `xml:"profiles"`
	Profiles    []Profile `xml:"profile,omitempty"`
	Func        string    `xml:"func,attr,omitempty"`
	Description string    `xml:"description,attr,omitempty"`
}
type Profile struct {
	XMLName     xml.Name  `xml:"profile"`
	Data        []Data    `xml:"data,omitempty"`
	Command     []Command `xml:"command,omitempty"`
	Description string    `xml:"description,attr,omitempty"`
	Value       string    `xml:"value,attr,omitempty"`
	Type        string    `xml:"type,attr,omitempty"`
}

type Data struct {
	XMLName xml.Name `xml:"data"`
	Enum    []Enum   `xml:"enum,omitempty"`
	Value   []Value  `xml:"value,omitempty"`
	Status  []Status `xml:"status,omitempty"`
	Command int      `xml:"command,omitempty"`
}

type Enum struct {
	XMLName     xml.Name    `xml:"enum"`
	Item        []Item      `xml:"item,omitempty"`
	RangeItem   []RangeItem `xml:"rangeitem,omitempty"`
	Description string      `xml:"description,attr,omitempty"`
	Shortcut    string      `xml:"shortcut,attr,omitempty"`
	Offset      int         `xml:"offset,attr,omitempty"`
	Size        int         `xml:"size,attr,omitempty"`
}

type Value struct {
	XMLName     xml.Name `xml:"value"`
	Range       Range    `xml:"range,omitempty"`
	Scale       Scale    `xml:"scale,omitempty"`
	Description string   `xml:"description,attr,omitempty"`
	Shortcut    string   `xml:"shortcut,attr,omitempty"`
	Offset      int      `xml:"offset,attr,omitempty"`
	Size        int      `xml:"size,attr,omitempty"`
	Unit        string   `xml:"unit,attr,omitempty"`
}

type Item struct {
	XMLName     xml.Name `xml:"item"`
	Description string   `xml:"description,attr,omitempty"`
	Value       string   `xml:"value,attr,omitempty"`
}
type RangeItem struct {
	XMLName     xml.Name `xml:"rangeitem"`
	Description string   `xml:"description,attr,omitempty"`
	Start       int      `xml:"start,attr"`
	End         *int     `xml:"end,attr"`
}
type Status struct {
	XMLName     xml.Name `xml:"status"`
	Description string   `xml:"description,attr,omitempty"`
	Shortcut    string   `xml:"shortcut,attr,omitempty"`
	Offset      int      `xml:"offset,attr,omitempty"`
	Size        int      `xml:"size,attr,omitempty"`
}

type Command struct {
	XMLName     xml.Name `xml:"command"`
	Item        string   `xml:"item,omitempty"`
	Description string   `xml:"description,attr,omitempty"`
	Shortcut    string   `xml:"shortcut,attr,omitempty"`
	Offset      int      `xml:"offset,attr,omitempty"`
	Size        int      `xml:"size,attr,omitempty"`
}

type Range struct {
	XMLName xml.Name `xml:"range"`
	Min     float64  `xml:"min,omitempty"`
	Max     float64  `xml:"max,omitempty"`
}
type Scale struct {
	XMLName xml.Name `xml:"scale"`
	Min     float64  `xml:"min,omitempty"`
	Max     float64  `xml:"max,omitempty"`
}

type EEP struct {
	Rorg      string
	Type      string
	Func      string
	Direction string
	Command   string
}

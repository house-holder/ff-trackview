package kml

import (
	"encoding/xml"
	"os"
)

type KML struct {
	XMLName  xml.Name  `xml:"kml"`
	Document *Document `xml:"Document"`
}

type Document struct {
	XMLName    xml.Name    `xml:"Document"`
	Placemarks []Placemark `xml:"Placemark"`
	// ExtendedData *ExtendedData `xml:"ExtendedData"`
}

type Placemark struct {
	Name     string `xml:"name"`
	StyleURL string `xml:"styleUrl"`
	Track    *Track `xml:"http://www.google.com/kml/ext/2.2 Track"`
}

type Track struct {
	AltitudeMode string   `xml:"altitudeMode"`
	Interpolate  string   `xml:"http://www.google.com/kml/ext/2.2 interpolate"`
	Coords       []Coord  `xml:"http://www.google.com/kml/ext/2.2 coord"`
	When         []string `xml:"when"`
}

type Coord struct {
	Text string `xml:",chardata"`
}

func ParseKML(filepath string) (*KML, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var kml KML
	err = xml.Unmarshal(data, &kml)
	return &kml, err
}

package metadata

import (
	"ff-trackview/internal/kml"
	"time"
)

type Metadata struct {
	Date        time.Time
	TailNumber  string
	FlightTitle string
}

func Extract(kml *kml.KML) (*Metadata, error) {
	// find <when>, parse date
	// find tailNumber & flightTitle
	md := &Metadata{ // NOTE: placeholder
		Date:        time.Now(),
		TailNumber:  "N123AB",
		FlightTitle: "KLAX - KJFK",
	}
	return md, nil
}

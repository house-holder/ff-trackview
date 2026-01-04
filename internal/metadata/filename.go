package metadata

import (
	"path/filepath"
	"regexp"
	"strings"
)

func GenerateFilename(m *Metadata, includeLocation bool) string {
	dateStr := m.Date.Format("2006-01-02")
	tail := sanitize(m.TailNumber)

	parts := []string{dateStr, tail}
	if includeLocation {
		flightTitle := sanitize(m.FlightTitle)
		parts = append(parts, flightTitle)
	}

	return strings.Join(parts, "-") + ".kml"
}

func sanitize(s string) string {
	s = strings.ReplaceAll(s, " - ", ",")
	s = strings.ReplaceAll(s, " ", "")
	re := regexp.MustCompile(`[<>:"/\\|?*]`)

	return re.ReplaceAllString(s, "_")
}

func FindAvailableFilename(dir, base string) (string, error) {
	fp := filepath.Base(base) // NOTE: placeholder
	// need to see if dup exists
	// try _1, _2, etc (first available)

	return fp, nil
}

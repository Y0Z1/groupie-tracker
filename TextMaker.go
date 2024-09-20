package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

func TextMaker(artist Bio) Text {
	var t Text

	// Determine if the artist is a solo or a band
	if len(artist.Members) == 1 {
		t.YearText = fmt.Sprintf("%s is a solo artist who began in %d.", artist.Name, artist.Year)
	} else {
		t.YearText = fmt.Sprintf("%s is a band formed in %d.", artist.Name, artist.Year)
	}
	// Add information about the first album
	t.AlbumText = fmt.Sprintf("Their first album was released on %s.", artist.Album)
	t.Members = artist.Members
	// Handle concerts sorting and formatting
	if len(artist.Concerts) > 0 {
		t.Atext = "Here are some of their concerts:"
		// Create a slice to hold the locations for sorting
		type locationInfo struct {
			location   string
			latestDate time.Time
		}
		var locations []locationInfo

		// Sort the concert dates and find the most recent date for each location
		dateFormat := "02-01-2006" // Date format for dd-mm-yyyy
		for location, dates := range artist.Concerts {
			sort.Slice(dates, func(i, j int) bool {
				dateI, err1 := time.Parse(dateFormat, dates[i])
				dateJ, err2 := time.Parse(dateFormat, dates[j])

				// Handle parsing errors by placing invalid dates at the end
				if err1 != nil || err2 != nil {
					return err1 == nil // If dateI parsed correctly but dateJ did not, keep dateI before dateJ
				}

				return dateI.After(dateJ) // Sort from newest to oldest
			})

			// Parse the most recent date
			latestDate, err := time.Parse(dateFormat, dates[0])
			if err == nil {
				locations = append(locations, locationInfo{location: location, latestDate: latestDate})
			}
		}

		// Sort locations by the most recent concert date in descending order
		sort.Slice(locations, func(i, j int) bool {
			return locations[i].latestDate.After(locations[j].latestDate)
		})

		// Build the paragraph with sorted locations and concert dates
		for _, loc := range locations {
			s := loc.location
			for i := 0; i < len(s); i++ {
				if s[i] == '-' {
					if len(s[i+2:]) < 3 {
						s = s[:i] + ", " + string(s[i+1]-32) + ToUpper(s[i+2:])
					} else {
						s = s[:i] + ", " + string(s[i+1]-32) + s[i+2:]
					}
				} else if s[i] == '_' {
					s = s[:i] + " " + s[i+1:]
				} else if i == 0 {
					s = string(s[0]-32) + s[1:]
				}
			}
			t.Concerts = append(t.Concerts, fmt.Sprintf("In %s:\t%s.", s, strings.Join(artist.Concerts[loc.location], ", ")))
		}
	} else {
		t.Atext = "Here are their albums:"
		t.Concerts = artist.Albums
	}
	return t
}

func ToUpper(s string) string {
	str := []rune(s)
	for i := 0; i < len(str); i++ {
		if str[i] >= 97 && str[i] <= 122 {
			str[i] = str[i] - 32
		}
	}
	ss := string(str)
	return ss
}

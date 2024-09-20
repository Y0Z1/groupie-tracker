package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Artist struct {
	ID        int      `json:"id"`
	Image     string   `json:"image"`
	Name      string   `json:"name"`
	Members   []string `json:"members"`
	Memb      string
	Memlen    string
	Albums    []string
	ABM       string
	Year      int    `json:"creationDate"`
	Album     string `json:"firstAlbum"`
	Loc       string `json:"locations"`
	Locations []string
	Con       string `json:"concertDates"`
	Rela      string `json:"relations"`
}

type Lcns struct{
	Locs []string `json:"locations"`
}
type index struct {
	L []Lcns `json:"index"`
}

type Event struct {
	ID       int                 `json:"id"`
	Relation map[string][]string `json:"datesLocations"`
}

type ArtistPage struct {
	Image  string
	Name   string
	Text   Text
	Memlen string
	Year   int
	Locations []string
}

type Text struct {
	YearText  string
	AlbumText string
	Members   []string
	Concerts  []string
	Albums    []string
	Atext     string
}

type Bio struct {
	Name     string
	Members  []string
	Year     int
	Album    string
	Albums   []string
	Concerts map[string][]string
}

func ParseHomeAPI() []Artist {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	as, err := http.Get(url)
	if err != nil {
		fmt.Print(err)
		return nil
	}
	defer as.Body.Close()
	data, err := io.ReadAll(as.Body)
	if err != nil {
		fmt.Print(err)
		return nil
	}
	url1 := "https://groupietrackers.herokuapp.com/api/locations"
	as1, err := http.Get(url1)
	if err != nil {
		fmt.Print(err)
		return nil
	}
	defer as.Body.Close()
	data1, err := io.ReadAll(as1.Body)
	if err != nil {
		fmt.Print(err)
		return nil
	}
	var ind index
	err = json.Unmarshal(data1, &ind)
	if err != nil {
		fmt.Print(err)
		return nil
	}
	var ar []Artist
	err = json.Unmarshal(data, &ar)
	cnt:=0
	for i, j := range ar {
		if j.ID!=0{
			for _, loc := range ind.L[cnt].Locs {
				s := loc
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
				ar[i].Locations = append(ar[i].Locations,s)
			}
			cnt++
		}
		if j.Memlen == "" {
			if len(j.Members) > 1 {
				ar[i].Memlen = fmt.Sprintf("%v Members", len(j.Members))
			} else {
				ar[i].Memlen = "Solo Artist"
			}
			for l, y := range j.Members {
				if l == 0 {
					ar[i].Memb += y
				} else {
					ar[i].Memb += "\n" + y
				}
			}
		}
		for l, y := range j.Albums {
			if l == 0 {
				ar[i].ABM += y
			} else {
				ar[i].ABM += "\n" + y
			}
		}
	}
	if err != nil {
		fmt.Print(err)
		return nil
	}

	return ar
}

func ParseArtistAPI(a Artist) ArtistPage {
	var AP ArtistPage
	var bio Bio
	url := a.Rela
	if url != "" {
		as, err := http.Get(url)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
		defer as.Body.Close()
		data, err := io.ReadAll(as.Body)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
		var api Event
		err = json.Unmarshal(data, &api)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
		bio.Concerts = api.Relation
	}
	AP.Image = a.Image
	AP.Memlen = a.Memlen
	AP.Name = a.Name
	AP.Year = a.Year
	bio.Name = a.Name
	bio.Album = a.Album
	bio.Year = a.Year
	bio.Albums = strings.Split(a.ABM, "\n")
	bio.Members = strings.Split(a.Memb, "\n")
	AP.Text = TextMaker(bio)
	AP.Locations = a.Locations
	return AP
}

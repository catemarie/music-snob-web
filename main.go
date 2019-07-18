package main

import (
    "fmt"
	"log"
    "net/http"
    "encoding/json"
    "strings"
)

type Results struct {
    Events   []string
}

var entries []string 

func render_entries(w http.ResponseWriter, r *http.Request){
    calendar_content := Results{
        Events: entries,
    }
    var jsonData []byte
    jsonData, err := json.Marshal(calendar_content)
    if err != nil {
        log.Println(err)
    }
    fmt.Fprintf(w, string(jsonData))
}

func main() {

    // get results for a city/state - San Diego, CA
    state := "California"
    cities := [3]string{ "San Diego", "Los Angeles", "Costa Mesa" }
    genre := "trance"

    for _, city := range cities {
        city_formatted := strings.Replace(city, " ", "%20", -1)
        log.Println(city)
        loc_id := get_location_id(state, city_formatted)
        artists := get_artists(loc_id)
        for _, artist  := range artists {
        	genres :=  get_artist_genres(artist.artists)   
    	    for _, g := range genres {
                if (g == genre || strings.Contains(g, genre)) {
                    line := artist.date + "    " + artist.artists
                    entries = append(entries, line)
                    break
                }
    	    }
        }
    }

    log.Println("entries: ", entries)
    http.HandleFunc("/results/", render_entries)
    http.ListenAndServe(":3000", nil)
}
package main

import (
    "fmt"
	"log"
    "net/http"
    "encoding/json"
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
    city := "San%20Diego"
    genre := "trance"

    loc_id := get_location_id(state, city)
    artists := get_artists(loc_id)
    for _, artist  := range artists {
    	genres :=  get_artist_genres(artist)   
	    for _, g := range genres {
            if (g == genre) {
                entries = append(entries, artist)
                break
            }
	    }
    }

    log.Println("entries: ", entries)

    http.HandleFunc("/results/", render_entries)
    http.ListenAndServe(":3000", nil)
}
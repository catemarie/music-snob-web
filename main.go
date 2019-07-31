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
    cities := [1]string{ "San Diego" }
    genre := "trance"

    init_database()

    for _, city := range cities {
        city_formatted := strings.Replace(city, " ", "%20", -1)

        loc_id := get_location_id(state, city_formatted)
        events := get_artists(loc_id)

        for _, event  := range events {

            var genres []string
            cached := read_database(event.artists)

            if len(cached) < 1 {
                genres =  get_artist_genres(event.artists)
                for _, new_genre := range genres {
                    write_database(event.artists, new_genre)
                }
            } else {
                genres = cached
            }

    	    for _, g := range genres {
                if (g == genre || strings.Contains(g, genre)) {
                    line := event.date + "    " + event.artists
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
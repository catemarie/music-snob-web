package main

import (
	"log"
)

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
                log.Println(artist)
                break
            }
	    }
    }
}
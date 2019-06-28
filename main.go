package main

import (
	"log"
)

func main() {

    // get results for a city/state - LA, CA
    state := "California"
    city := "San%20Diego"

    loc_id := get_location_id(state, city)
    artists := get_artists(loc_id)
    for _, artist  := range artists {
    	log.Println(artist)
    	genres :=  get_artist_genres(artist)   
	    for _, genre := range genres {
	        log.Println("    * " + genre)
	    }
    }
}
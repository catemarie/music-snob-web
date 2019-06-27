package main

func main() {

    // get results for a city/state - LA, CA
    state := "California"
    city := "San%20Diego"

    loc_id := get_location_id(state, city)
    get_artists(loc_id)
}
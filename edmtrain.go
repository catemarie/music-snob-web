package main

import (
    "os"
    "net/http"
    "log"
    "io/ioutil"
    "github.com/tidwall/gjson"
)

var key = os.ExpandEnv("$EDMTRAIN_KEY")

func get_json(url string) string {

    // API call
    resp, err := http.Get(url)
    if err != nil {
        log.Println("Couldn't fetch results")
    }

    // grab the JSON
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Println("Couldn't read results")
    }

    return string(body)

}
func get_location_id(state string, city string) string {
    
    url := "https://edmtrain.com/api/locations?state=" + state + "&city=" + city + "&client=" + key
    body := get_json(url)

    // get the location ID
    value := gjson.Get(string(body), "data.#.id")
    loc_id := value.Array()[0].String()

    return loc_id
}

func get_artists(loc_id string) []string {

    var a []string

    url := "https://edmtrain.com/api/events?locationIds=" + loc_id + "&client=" + key
    body := get_json(url)

    value := gjson.Get(string(body), "data.#.artistList.#.name")
    next := value.Array()

    for _, n := range next {
        artist := n.Array()
        if len(artist) > 0 {
            a = append(a, artist[0].String())
        }
    }

    return a
}

package main

import (
    "os"
    "net/http"
    "log"
    "io/ioutil"
    "github.com/tidwall/gjson"
    "strconv"
)

type Event struct {
    date string
    artists string
}

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

func get_artists(loc_id string) []Event {

    var a []Event

    url := "https://edmtrain.com/api/events?locationIds=" + loc_id + "&client=" + key
    body := get_json(url)

    json_txt := gjson.Get(string(body), "data")
    test := json_txt.Array()
    num := len(test)

    for ii := 0; ii < num; ii++ {

        date := gjson.Get(string(body), "data." + strconv.Itoa(ii) + ".date")
        artistList := gjson.Get(string(body), "data." + strconv.Itoa(ii) + ".artistList.#.name")

        for _, n := range artistList.Array() {
            artist := n.Array()
            if len(artist) > 0 {
                entry := Event{date.String(), artist[0].String()}
                a = append(a, entry)
            }
        }
    }

    return a
}

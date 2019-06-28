package main

import (
    "context"
    "log"
    "os"
    "golang.org/x/oauth2/clientcredentials"
    "github.com/zmb3/spotify"
)

var client_id = os.ExpandEnv("$SPOTIFY_CLIENT_ID")
var client_secret = os.ExpandEnv("$SPOTIFY_CLIENT_SECRET")

func get_artist_genres(artist string) []string {

    // handle credentials
    config := &clientcredentials.Config{
        ClientID:     client_id,
        ClientSecret: client_secret,
        TokenURL:     spotify.TokenURL,
    }
    token, err := config.Token(context.Background())
    if err != nil {
        log.Fatalf("couldn't get token: %v", err)
    }
    client := spotify.Authenticator{}.NewClient(token)

    // search for this artist and get the top result
    results, err := client.Search(artist, spotify.SearchTypeArtist)
    if err != nil {
        log.Fatal(err)
    }

    // list all of the artist's genres
    var g []string
    if results.Artists != nil {
        artists := results.Artists.Artists
        if len(artists) > 0 {
            for _, genre := range artists[0].Genres {
                g = append(g, genre)
            }
        }
    }
    
    return g
}

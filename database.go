package main

import (
    "log"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init_database() {

    var err error

    db, err = sql.Open("sqlite3", "./ms.db")
    if err != nil {
        db.Close()
        log.Fatal(err.Error())
    }

    // create database
    var schema = `
    CREATE TABLE IF NOT EXISTS Events (
        id integer PRIMARY KEY AUTOINCREMENT,
        artist text,
        genre text
    );
    `

    _, err = db.Exec(schema)
        if err != nil {
        db.Close()
        log.Fatal(err.Error())
    }

    db.Close()
}

func write_database(artist string, genre string) {

    var err error
    
    db, err = sql.Open("sqlite3", "./ms.db")
    if err != nil {
        db.Close()
        log.Fatal(err.Error())
    }

    query := `
    INSERT INTO Events(artist, genre)
    VALUES(?, ?)
    `

    stmt, err := db.Prepare(query)
    if err != nil {
        db.Close()
        log.Fatal(err.Error())
    }

    _, err = stmt.Exec(artist, genre)
    if err != nil {
        db.Close()
        log.Fatal(err.Error())
    }

    db.Close()
}

func read_database(artist string) []string {

    var err error

    db, err = sql.Open("sqlite3", "./ms.db")
    if err != nil {
        db.Close()
        log.Fatal(err.Error())
    }

    query := `
    SELECT * 
    FROM Events
    WHERE artist=?;
    `

    stmt, err := db.Prepare(query)
    if err != nil {
        db.Close()
        log.Fatal(err.Error())
    }

    rows, err := stmt.Query(artist)
    if err != nil {
        db.Close()
        log.Fatal(err.Error())
    }

    var id int
    var genres []string
    for rows.Next() {
        var genre string
        err = rows.Scan(&id, &artist, &genre)
        if err != nil {
            db.Close()
            log.Fatal(err.Error())
        }
        genres = append(genres, genre)
    }
    rows.Close()

    db.Close()

    return genres
}

func main() {

    init_database()

    genres := read_database("andrew rayel")
    for _, genre := range genres {
        log.Printf(genre)
    }
    if len(genres) < 1{
        write_database("andrew rayel", "trance")
    }
}
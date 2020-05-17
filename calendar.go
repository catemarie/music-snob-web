package main

import (
	"github.com/lestrrat/go-ical"
	"os"
)

func main() {

	c := ical.New()
	c.AddProperty("X-Foo-Bar-Baz", "value")
	tz := ical.NewTimezone()
	tz.AddProperty("TZID", "Asia/Tokyo")
	c.AddEntry(tz)

	ical.NewEncoder(os.Stdout).Encode(c)
}

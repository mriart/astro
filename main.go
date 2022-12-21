package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/zsefvlol/timezonemapper"
)

const preHTML = `
<html>
<head>
    <meta charset='UTF-8'>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <style>
    body {
        background-color: black;
        color: white;
    }
    </style>
</head>

<body>
    <pre>
`

const postHTML = `
    </pre>
</body>
</html>
`

func astro(w http.ResponseWriter, r *http.Request) {
	// Build the response
	var resp string = ""

	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	// Evaluate if the URL has the location paramenters
	q := r.URL.Query()
	if len(q) == 0 {
		http.ServeFile(w, r, "getLatLon.html")
		return
	}

	// URL has coordinates (lat & lon), get them, convert to float
	lat, _ := strconv.ParseFloat(q["lat"][0], 64)
	lon, _ := strconv.ParseFloat(q["lon"][0], 64)
	fmt.Println(lat, lon)

	// Get the time now for the observation, and time zone from lat & lon
	t := time.Now()
	tz := timezonemapper.LatLngToTimezoneString(lat, lon)
	loc, err := time.LoadLocation(tz)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Timezone: %s\nTime in location: %s\n", tz, t.In(loc))

	// Load date, time and coordinates in responses
	resp += fmt.Sprintf("Date, time: %s\n", t.In(loc).Round(time.Second))
	resp += fmt.Sprintf("Timezone location: %s\n", loc)
	resp += fmt.Sprintf("Latitude: %f°\nLongitude: %f°\n\n", lat, lon)

	// Load sun, moon & planets data
	resp += getSunData(lat, lon, t, loc)
	resp += getMoonData(lat, lon, t, loc)
	x := getPlanetsData(lat, lon, t, loc)
	resp += x

	fmt.Printf("Petition received:\n%f\n%f\n%s\n%s\n", lat, lon, t, loc)
	//fmt.Println(resp)
	fmt.Fprintf(w, "%s", preHTML+resp+postHTML)
}

func main() {
	http.HandleFunc("/", astro)

	fmt.Printf("Starting server, port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

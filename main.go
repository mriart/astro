package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

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
	// Idem for time now
	lat, _ := strconv.ParseFloat(q["lat"][0], 64)
	lon, _ := strconv.ParseFloat(q["lon"][0], 64)
	t := time.Now()

	// Load date, time and coordinates in responses
	resp += fmt.Sprintf("Date, time: %s\n", t.Round(time.Second))
	resp += fmt.Sprintf("Latitude: %f\nLongitude: %f\n\n", lat, lon)

	// Load sun, moon & planets data
	resp += getSunData(lat, lon, t)
	resp += getMoonData(lat, lon, t)
	x := getPlanetsData(lat, lon, t)
	resp += x

	//fmt.Println(resp)
	fmt.Fprintf(w, "%s", resp)
}

func main() {
	http.HandleFunc("/", astro)

	fmt.Printf("Starting server, port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

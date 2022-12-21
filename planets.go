package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/tidwall/gjson"
)

const RapidAPIHost = "astronomy.p.rapidapi.com"

type Planet struct {
	Name      string
	Altitude  float64
	Azimuth   float64
	Time      string
	Magnitude float64
}

// Method to fill the struct Planet with the data provided from the API response
func (p *Planet) GetCoordinates(planetName string, pData *string) {
	idx := ""
	switch planetName {
	case "Mercury":
		idx = "2"
	case "Venus":
		idx = "3"
	case "Mars":
		idx = "5"
	case "Jupiter":
		idx = "6"
	case "Saturn":
		idx = "7"
	}
	p.Name = gjson.Get(*pData, "data.table.rows."+idx+".cells.0.name").Str
	p.Altitude = gjson.Get(*pData, "data.table.rows."+idx+".cells.0.position.horizontal.altitude.degrees").Float()
	p.Azimuth = gjson.Get(*pData, "data.table.rows."+idx+".cells.0.position.horizontal.azimuth.degrees").Float()
	p.Time = time.Now().String()
	p.Magnitude = gjson.Get(*pData, "data.table.rows."+idx+".cells.0.extraInfo.magnitude").Float()
}

// Get the data from the API provider, RapidAPI with the url and credentials above (in const).
// Returns a json in format string to be processed by the library gjson
func getAPIData(lat float64, lon float64, t time.Time, loc *time.Location) string {
	url := fmt.Sprintf(
		"%s%s%s%f%s%f%s%s%s%s%s%s%s",
		"https://",
		RapidAPIHost,
		"/api/v2/bodies/positions?latitude=", lat,
		"&longitude=", lon,
		"&from_date=", t.Format("2006-01-02"),
		"&to_date=", t.Format("2006-01-02"),
		"&elevation=1",
		"&time=", url.QueryEscape(t.In(loc).Format("15:04:05")),
	)
	//fmt.Println(url)
	//fmt.Println(t)

	RapidAPIKey := os.Getenv("RAPID_API_KEY")
	if RapidAPIKey == "" {
		fmt.Println("Missing RAPID_API_KEY env variable")
	}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-RapidAPI-Key", RapidAPIKey)
	req.Header.Add("X-RapidAPI-Host", RapidAPIHost)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	//fmt.Println(string(body))
	return string(body)
}

func getPlanetsData(lat float64, lon float64, t time.Time, loc *time.Location) string {
	// Variable to load the response and return
	var resp string = ""

	// Get the data from the API provider and load into apiData string
	apiData := getAPIData(lat, lon, t, loc)

	// Declare, initialize and assing Planet struct for each planet
	mercury := Planet{}
	mercury.GetCoordinates("Mercury", &apiData)

	venus := Planet{}
	venus.GetCoordinates("Venus", &apiData)

	mars := Planet{}
	mars.GetCoordinates("Mars", &apiData)

	jupiter := Planet{}
	jupiter.GetCoordinates("Jupiter", &apiData)

	saturn := Planet{}
	saturn.GetCoordinates("Saturn", &apiData)

	// Sprintf and return
	resp += fmt.Sprintf("*Mercury\nAltitude: %.2f°\nAzimuth(from N): %.2f°\nMagnitude: %.2f\n\n", mercury.Altitude, mercury.Azimuth, mercury.Magnitude)
	resp += fmt.Sprintf("*Venus\nAltitude: %.2f°\nAzimuth(from N): %.2f°\nMagnitude: %.2f\n\n", venus.Altitude, venus.Azimuth, venus.Magnitude)
	resp += fmt.Sprintf("*Mars\nAltitude: %.2f°\nAzimuth(from N): %.2f°\nMagnitude: %.2f\n\n", mars.Altitude, mars.Azimuth, mars.Magnitude)
	resp += fmt.Sprintf("*Jupiter\nAltitude: %.2f°\nAzimuth(from N): %.2f°\nMagnitude: %.2f\n\n", jupiter.Altitude, jupiter.Azimuth, jupiter.Magnitude)
	resp += fmt.Sprintf("*Saturn\nAltitude: %.2f°\nAzimuth(from N): %.2f°\nMagnitude: %.2f\n\n", saturn.Altitude, saturn.Azimuth, saturn.Magnitude)

	//fmt.Println(resp)
	return resp
}

package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/tidwall/gjson"
)

const urlPre = "https://astronomy.p.rapidapi.com"
const RapidAPIKey = "d290331181msh2e6a23ba8aff482p1cd1b9jsn491faa927e0b"
const RapidAPIHost = "astronomy.p.rapidapi.com"

type Planet struct {
	Name     string
	Altitude float64
	Azimuth  float64
	Time     string
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
}

// Get the data from the API provider, RapidAPI with the url and credentials above (in const).
// Returns a json in format string to be processed by the library gjson
func getAPIData(lat float64, lon float64, t time.Time) string {
	url := fmt.Sprintf(
		"%s%s%f%s%f%s%s%s%s%s%s%s",
		urlPre,
		"/api/v2/bodies/positions?latitude=", lat,
		"&longitude=", lon,
		"&from_date=", t.Format("2006-01-02"),
		"&to_date=", t.Format("2006-01-02"),
		"&elevation=1",
		"&time=", t.Format("15%3A04%3A05"),
	)
	fmt.Println(url)
	fmt.Println(t)

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

func getPlanetsData(lat float64, lon float64, t time.Time) string {
	// Variable to load the response and return
	var resp string = ""

	// Get the data from the API provider and load into apiData string
	apiData := getAPIData(lat, lon, t)

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
	resp += fmt.Sprintf("*Mercury\nAltitude: %.2f°\nAzimuth(from N): %.2f°\n\n", mercury.Altitude, mercury.Azimuth)
	resp += fmt.Sprintf("*Venus\nAltitude: %.2f°\nAzimuth(from N): %.2f°\n\n", venus.Altitude, venus.Azimuth)
	resp += fmt.Sprintf("*Mars\nAltitude: %.2f°\nAzimuth(from N): %.2f°\n\n", mars.Altitude, mars.Azimuth)
	resp += fmt.Sprintf("*Jupiter\nAltitude: %.2f°\nAzimuth(from N): %.2f°\n\n", jupiter.Altitude, jupiter.Azimuth)
	resp += fmt.Sprintf("*Saturn\nAltitude: %.2f°\nAzimuth(from N): %.2f°\n\n", saturn.Altitude, saturn.Azimuth)

	//fmt.Println(resp)
	return resp
}

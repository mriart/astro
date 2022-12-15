package main

import (
	"fmt"
	"time"

	"github.com/sixdouglas/suncalc"
)

func getSunData(lat float64, lon float64, t time.Time) string {
	// Variable to load the response and return
	var resp string = ""

	// Get sun data lat & lon in float
	times := suncalc.GetTimes(t, lat, lon)

	sRise := times[suncalc.Sunrise].Time.Round(time.Second)
	sSet := times[suncalc.Sunset].Time.Round(time.Second)
	sDuration := sSet.Sub(sRise)

	resp += "*Sun\n"
	resp += fmt.Sprintf("Sunrise: %s\n", sRise)
	resp += fmt.Sprintf("Sunset: %s\n", sSet)
	resp += fmt.Sprintf("Day duration: %s\n\n", sDuration)

	//fmt.Prinln(resp)
	return resp
}

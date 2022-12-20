package main

import (
	"fmt"
	"time"

	"github.com/sixdouglas/suncalc"
)

// This function returns the day duration difference since yesterday. There are 3 returns:
// 1) Day duration diff (if negative), the day is shorter
// 2) Diff of the sunrising
// 3) Diff of the sunsetting
func diffDayDuration(t time.Time, lat float64, lon float64) (time.Duration, time.Duration, time.Duration) {
	times := suncalc.GetTimes(t, lat, lon)
	sRise := times[suncalc.Sunrise].Time
	sSet := times[suncalc.Sunset].Time
	fmt.Printf("Today, rise: %s, set: %s\n", sRise, sSet)

	tYest := t.Add(-24 * time.Hour)
	timesYest := suncalc.GetTimes(tYest, lat, lon)
	sRiseYest := timesYest[suncalc.Sunrise].Time
	sSetYest := timesYest[suncalc.Sunset].Time
	//fmt.Printf("Yestd, rise: %s, set: %s\n", sRiseYest, sSetYest)

	diffRise := sRiseYest.Add(24 * time.Hour).Sub(sRise)
	diffSet := -sSetYest.Add(24 * time.Hour).Sub(sSet)
	diffDay := diffRise + diffSet
	//fmt.Printf("Diff:, rise: %s, set: %s\n", diffRise, diffSet)

	return diffRise.Round(time.Second), diffSet.Round(time.Second), diffDay.Round(time.Second)
}

// Get data for Sun and return string to main
func getSunData(lat float64, lon float64, t time.Time) string {
	// Variable to load the response and return
	var resp string = ""

	// Get sun data lat & lon in float
	times := suncalc.GetTimes(t, lat, lon)

	sRise := times[suncalc.Sunrise].Time.Round(time.Second)
	sSet := times[suncalc.Sunset].Time.Round(time.Second)
	sDuration := sSet.Sub(sRise)
	diffRise, diffSet, diffDay := diffDayDuration(t, lat, lon)

	resp += "*Sun\n"
	resp += fmt.Sprintf("Sunrise: %s\n", sRise)
	resp += fmt.Sprintf("Sunset: %s\n", sSet)
	resp += fmt.Sprintf("Day duration: %s\n", sDuration)
	resp += fmt.Sprintf("Diff since yesterday: %s (%s/%s)\n\n", diffDay, diffRise, diffSet)

	//fmt.Prinln(resp)
	return resp
}

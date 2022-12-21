package main

import (
	"fmt"
	"time"

	"github.com/IvanMenshykov/MoonPhase"
)

func getMoonData(lat float64, lon float64, t time.Time, loc *time.Location) string {
	// Variable to load the response and return
	var resp string = ""

	// Declar a MoonPhase var, initialize and assign
	m := MoonPhase.New(t)

	resp += "*Moon\n"
	resp += fmt.Sprintf("Age: %f days\n", m.Age())
	resp += fmt.Sprintf("Phase (0-1): %f\n", m.Phase())
	resp += fmt.Sprintf("Illumination (0-1): %f\n", m.Illumination())

	if m.Phase() < 0.5 {
		fullMoonTime := time.Unix(int64(m.FullMoon()), 0)
		resp += fmt.Sprintf("Full moon: %s\n", fullMoonTime.In(loc))
		newMoonTime := time.Unix(int64(m.NextNewMoon()), 0)
		resp += fmt.Sprintf("New moon: %s\n", newMoonTime.In(loc))
	} else {
		newMoonTime := time.Unix(int64(m.NextNewMoon()), 0)
		resp += fmt.Sprintf("New moon: %s\n", newMoonTime.In(loc))
		fullMoonTime := time.Unix(int64(m.NextFullMoon()), 0)
		resp += fmt.Sprintf("Full moon: %s\n", fullMoonTime.In(loc))
	}

	resp += fmt.Sprintf("Zodiac: %s\n\n", m.ZodiacSign())

	//fmt.Println(resp)
	return resp
}

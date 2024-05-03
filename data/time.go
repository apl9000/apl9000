package data

import (
	"time"
)

func GetCurrentTime() string {
	currentTime := time.Now()
	location, _ := time.LoadLocation("America/New_York")
  timeInNY := currentTime.In(location)
	return timeInNY.Format("2006-01-02 15:04:05")
}
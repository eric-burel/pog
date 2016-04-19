package timetogo

import "time"

func duration(i int) (d time.Duration) {
	d = time.Duration(i)
	return
}

// Seconds Converts an value value in minutes
func Seconds(i int) (d time.Duration) {
	d = duration(i) * time.Second
	return
}

// Minutes Converts an value value in minutes
func Minutes(i int) (d time.Duration) {
	d = duration(i) * time.Minute
	return
}

// Hours Converts an value value in minutes
func Hours(i int) (d time.Duration) {
	d = duration(i) * time.Hour
	return
}

// Days Converts an value value in minutes
func Days(i int) (d time.Duration) {
	d = duration(i) * time.Hour * 24
	return
}

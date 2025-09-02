package utils

import "time"

func LocJakarta() *time.Location {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	return loc
}
func TimeNowJkrt() time.Time {
	now := time.Now().In(LocJakarta())
	return now
}

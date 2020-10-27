package model

import "time"

type CaptorValue struct {
	Id         int       `example:325936881483120640`
	CaptorId   int       `example:4`
	AirportId  string    `example:"BDX"`
	Type       string    `example:"Temperature"`
	Date       time.Time `example:"2020-10-27T01:52:00Z"`
	StringDate string    `example:"20201027"`
	Value      float64   `example:14.75"`
}

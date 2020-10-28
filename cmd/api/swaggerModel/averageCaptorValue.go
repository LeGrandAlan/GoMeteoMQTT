package swaggerModel

import "time"

type AverageCaptorValue struct {
	AirportId   string    `example:"BDX"`
	Date        time.Time `example:"2020-10-27"`
	Temperature float64   `example:"14.56"`
	Pressure    float64   `example:"16.86"`
	Wind        float64   `example:"13.56"`
}

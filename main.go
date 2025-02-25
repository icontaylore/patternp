package main

import (
	"headFirst/Observer_pattern/observer/obs1"
	"headFirst/Observer_pattern/weather_data"
)

func main() {
	weather := weather_data.Weather{}
	subscriberOne := obs1.ClientOne{Name: "Клиент - 1"}

	weather.RegisterObs(&subscriberOne)
	weather.SetMeasurements(22, 22, 22)
	weather.GetInfo()
}

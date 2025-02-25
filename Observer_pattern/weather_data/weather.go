package weather_data

import (
	"fmt"
	"headFirst/Observer_pattern/observer"
)

type Weather struct {
	Array []observer.Observer
	temp  float64
	humid float64
	press float64
}

func (w *Weather) GetInfo() {
	fmt.Println(w.humid, w.temp, w.press)
}

func (w *Weather) RegisterObs(client observer.Observer) {
	w.Array = append(w.Array, client)
}

func (w *Weather) DeleteObs(client observer.Observer) {

}

func (w *Weather) NotifyObs() {
	for _, v := range w.Array {
		v.Update(w.temp, w.press, w.humid)
	}
}

func (w *Weather) SetMeasurements(temp, humid, press float64) {
	w.temp = temp
	w.humid = humid
	w.press = press
	w.NotifyObs()
}

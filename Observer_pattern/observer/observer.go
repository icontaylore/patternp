package observer

type Observer interface {
	Update(temp, humid, press float64)
}

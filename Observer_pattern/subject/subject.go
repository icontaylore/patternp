package subject

import "headFirst/Observer_pattern/observer"

type Subject interface {
	RegisterObs(client observer.Observer)
	DeleteObs(client observer.Observer)
	NotifyObs()
}

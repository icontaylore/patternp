package main

import (
	"headFirst/Strategy pattern/Duck"
)

func main() {

	var duck Duck.Ducker
	duck = &Duck.RezinDuck{Name: "Утка Lame", FlyBehavior: &Duck.FlyNot{}}
	duck.ExclusiveFly()

	duck = &Duck.RezinDuck{Name: "Утка John", FlyBehavior: &Duck.FlyWithWings{}}
	duck.ExclusiveFly()

}

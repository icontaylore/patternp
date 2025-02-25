package main

import (
	"headFirst/Strategy_pattern/Duck"
)

func main() {

	var duck Duck.Ducker
	duck = &Duck.DefaultDuck{Name: "Утка Lame", FlyBehavior: &Duck.FlyNot{}}
	duck.ExclusiveFly()

	duck = &Duck.RezinDuck{Name: "Утка John", FlyBehavior: &Duck.FlyWithWings{}}
	duck.ExclusiveFly()

	duck = &Duck.PrimankaDuck{Name: "Игрушечная утка", FlyBehavior: &Duck.FlyNot{}}
	duck.ExclusiveFly()
	duck.SetFly(&Duck.FlyWithRocket{})
	duck.ExclusiveFly()
}

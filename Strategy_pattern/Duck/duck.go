package Duck

import "fmt"

type FlyBehavior interface {
	Fly(s string)
}

type Ducker interface {
	ExclusiveFly()
	SetFly(fb FlyBehavior)
}

// с ракетой
type FlyWithRocket struct {
}

// с крыльями
type FlyWithWings struct {
}

// не умеет летать
type FlyNot struct {
}

func (f *FlyWithWings) Fly(name string) {
	fmt.Println("могу летать", name)
}

func (F *FlyWithRocket) Fly(name string) {
	fmt.Println("Летит с реактивным двигалетем", name)
}

func (f *FlyNot) Fly(name string) {
	fmt.Println("не могу летать", name)
}

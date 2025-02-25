package Duck

import "fmt"

type FlyBehavior interface {
	Fly(s string)
}

type Ducker interface {
	ExclusiveFly()
}

type FlyWithWings struct {
}

type FlyNot struct {
}

type RezinDuck struct {
	Name        string
	FlyBehavior FlyBehavior
}

type DefaultDuck struct {
	Name        string
	FlyBehavior FlyBehavior
}

func (f *FlyWithWings) Fly(name string) {
	fmt.Println("могу летать", name)
}

func (f *FlyNot) Fly(name string) {
	fmt.Println("не могу летать", name)
}

func (r *RezinDuck) ExclusiveFly() {
	r.FlyBehavior.Fly(r.Name)
}

func (d *DefaultDuck) ExclusiveFly() {
	d.FlyBehavior.Fly(d.Name)
}

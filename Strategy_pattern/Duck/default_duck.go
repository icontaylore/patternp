package Duck

type DefaultDuck struct {
	Name        string
	FlyBehavior FlyBehavior
}

// Default Duck
func (d *DefaultDuck) ExclusiveFly() {
	d.FlyBehavior.Fly(d.Name)
}
func (d *DefaultDuck) SetFly(fb FlyBehavior) {
	d.FlyBehavior = fb
}

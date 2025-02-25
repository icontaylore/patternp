package Duck

type PrimankaDuck struct {
	Name        string
	FlyBehavior FlyBehavior
}

// Primanka Duck
func (d *PrimankaDuck) ExclusiveFly() {
	d.FlyBehavior.Fly(d.Name)
}
func (d *PrimankaDuck) SetFly(fb FlyBehavior) {
	d.FlyBehavior = fb
}

//

package Duck

type RezinDuck struct {
	Name        string
	FlyBehavior FlyBehavior
}

// Rezin Duck
func (r *RezinDuck) ExclusiveFly() {
	r.FlyBehavior.Fly(r.Name)
}
func (r *RezinDuck) SetFly(fb FlyBehavior) {
	r.FlyBehavior = fb
}

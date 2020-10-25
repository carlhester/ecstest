package components

type DamageComponent struct {
	Damage int
}

func (c DamageComponent) IsComponent() bool {
	return true
}

package components

type HealthComponent struct {
	Hp int
}

func (c HealthComponent) IsComponent() bool {
	return true
}

package components

type HealComponent struct {
	Hp int
}

func (c HealComponent) IsComponent() bool {
	return true
}

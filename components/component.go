package components

type Component interface {
	IsComponent() bool
}

type HealthComponent struct {
	Hp int
}

func (c HealthComponent) IsComponent() bool {
	return true
}

type HealComponent struct {
	Hp int
}

func (c HealComponent) IsComponent() bool {
	return true
}

type DamageComponent struct {
	Damage int
}

func (c DamageComponent) IsComponent() bool {
	return true
}

type SizeComponent struct {
    Height int
}

func (c SizeComponent) IsComponent() bool {
	return true
}

type GrowComponent struct { 
    Growth int
}

func (c GrowComponent) IsComponent() bool {
	return true
}

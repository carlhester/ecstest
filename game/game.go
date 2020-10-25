package game


import (
	"ecstest/entity"
	"ecstest/components"
	"ecstest/systems"
	"fmt"
)

type game struct {}

func NewGame() *game{
    return &game{}
}

func (g *game)Run() { 
	var entities []*entity.Entity

	e := entity.NewEntity()
	e.AddComponent("health", components.HealthComponent{Hp: 100})
	e.AddComponent("heal", components.HealComponent{Hp: 10})
	e.AddComponent("damage", components.DamageComponent{Damage: 3})
	entities = append(entities, e)

	for i, e := range entities {
		fmt.Printf("id: %d. %+v\n\n", i, e)
	}

	for _, e := range entities {
		systems.HealEntities(e)
		systems.DamageEntities(e)
	}

	for i, e := range entities {
		fmt.Printf("id: %d. %+v\n\n", i, e)
	}
}

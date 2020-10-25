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

	player := entity.NewEntity()
	player.AddComponent("health", components.HealthComponent{Hp: 100})
	player.AddComponent("heal", components.HealComponent{Hp: 10})
    player.AddComponent("damage", components.DamageComponent{Damage: 3})
	entities = append(entities, player)

    tree := entity.NewEntity()
	tree.AddComponent("size", components.SizeComponent{Height: 125})
	tree.AddComponent("grow", components.GrowComponent{Growth: 8})
	entities = append(entities, tree)


	for i, e := range entities {
		fmt.Printf("id: %d. %+v\n", i, e)
	}

    fmt.Printf("\n\nupdates\n\n")
	for _, e := range entities {
		systems.HealEntities(e)
		systems.DamageEntities(e)
        systems.GrowEntities(e)
	}

	for i, e := range entities {
		fmt.Printf("id: %d. %+v\n", i, e)
	}
}

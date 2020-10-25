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

	player := entity.NewEntity("Player")
	player.AddComponent("health", components.HealthComponent{Hp: 100})
	player.AddComponent("heal", components.HealComponent{Hp: 10})
    player.AddComponent("damage", components.DamageComponent{Damage: 3})
	entities = append(entities, player)
	
    mob := entity.NewEntity("Mob")
	mob.AddComponent("health", components.HealthComponent{Hp: 100})
    mob.AddComponent("damage", components.DamageComponent{Damage: 17})
	entities = append(entities, mob)

    tree := entity.NewEntity("Tree")
	tree.AddComponent("size", components.SizeComponent{Height: 125})
	tree.AddComponent("grow", components.GrowComponent{Growth: 1})
	entities = append(entities, tree)


	for i, e := range entities {
		fmt.Printf("%s\tid: %d\t%+v\n", e.Name, i, e)
	}

    fmt.Printf("\nUpdating Systems Start.\n\n")
	for _, e := range entities {
		systems.HealEntities(e)
		systems.DamageEntities(e)
        systems.GrowEntities(e)
	}
    fmt.Printf("\nUpdating Systems Complete.\n\n")

	for i, e := range entities {
		fmt.Printf("%s\tid: %d\t%+v\n", e.Name, i, e)
	}
}

package systems

import (
	"ecstest/components"
	"ecstest/entity"
)


func HealEntities(e *entity.Entity) {
	if e.HasComponent("heal") {
		if e.HasComponent("health") {
			hp := e.Components["health"].(components.HealthComponent).Hp + e.Components["heal"].(components.HealComponent).Hp
			e.RemoveComponent("health")
			e.RemoveComponent("heal")
			e.AddComponent("health", components.HealthComponent{Hp: hp})
		}
	}
}

func DamageEntities(e *entity.Entity) {
	if e.HasComponent("damage") {
		hp := e.Components["health"].(components.HealthComponent).Hp - e.Components["damage"].(components.DamageComponent).Damage
		e.RemoveComponent("health")
		e.RemoveComponent("damage")
		e.AddComponent("health", components.HealthComponent{Hp: hp})
	}
}

func GrowEntities(e *entity.Entity) {
	if e.HasComponent("size") {
	    if e.HasComponent("grow") {
            height := e.Components["size"].(components.SizeComponent).Height + e.Components["grow"].(components.GrowComponent).Growth
            e.RemoveComponent("size")
            e.RemoveComponent("grow")
            e.AddComponent("size", components.SizeComponent{Height: height})
	    }
	}
}



package systems

import (
	"ecstest/components"
	"ecstest/entity"
	"fmt"
)


func HealEntities(e *entity.Entity) {
	fmt.Println("healing ", e)
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
	fmt.Println("damaging", e)
	if e.HasComponent("damage") {
		hp := e.Components["health"].(components.HealthComponent).Hp - e.Components["damage"].(components.DamageComponent).Damage
		e.RemoveComponent("health")
		e.RemoveComponent("damage")
		e.AddComponent("health", components.HealthComponent{Hp: hp})
	}
}


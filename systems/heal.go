package systems

import (
    "fmt"
	"ecstest/components"
	"ecstest/entity"
)


func HealEntities(e *entity.Entity) {
	if e.HasComponent("heal") {
		if e.HasComponent("health") {
            fmt.Printf("HEAL: %s heals for %d\n", e.Name, e.Components["heal"].(components.HealComponent).Hp)
			hp := e.Components["health"].(components.HealthComponent).Hp + e.Components["heal"].(components.HealComponent).Hp
			e.RemoveComponent("health")
			e.RemoveComponent("heal")
			e.AddComponent("health", components.HealthComponent{Hp: hp})
		}
	}
}

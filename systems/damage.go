package systems

import (
	"fmt"

	"github.com/crucialcarl/ecstest/components"
	"github.com/crucialcarl/ecstest/entity"
)

func DamageEntities(e *entity.Entity) {
	if e.HasComponent("damage") {
		fmt.Printf("DAMAGE: %s is damaged for %d\n", e.Name, e.Components["damage"].(components.DamageComponent).Damage)
		hp := e.Components["health"].(components.HealthComponent).Hp - e.Components["damage"].(components.DamageComponent).Damage
		e.RemoveComponent("health")
		e.RemoveComponent("damage")
		e.AddComponent("health", components.HealthComponent{Hp: hp})
	}
}

package systems

import (
    "fmt"
	"ecstest/components"
	"ecstest/entity"
)

func GrowEntities(e *entity.Entity) {
	if e.HasComponent("size") {
	    if e.HasComponent("grow") {
            fmt.Printf("GROW: %s grows by %d\n", e.Name, e.Components["grow"].(components.GrowComponent).Growth)
            height := e.Components["size"].(components.SizeComponent).Height + e.Components["grow"].(components.GrowComponent).Growth
            e.RemoveComponent("size")
            e.RemoveComponent("grow")
            e.AddComponent("size", components.SizeComponent{Height: height})
	    }
	}
}



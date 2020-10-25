package entity

import (
	"ecstest/components"
)

type Entity struct {
    Name string
	Components map[string]components.Component
}

func NewEntity(name ...string) *Entity {
    var charName string
    if len(name) == 0 { 
        charName = "UnnamedEntity"
    } else { 
        charName = name[0]
    }

	return &Entity{
        Name: charName,
		Components: make(map[string]components.Component),
	}
}

func (e *Entity) HasComponent(componentName string) bool {
	if _, ok := e.Components[componentName]; ok {
		return true
	}
	return false
}

func (e *Entity) AddComponent(componentName string, component components.Component) {
	e.Components[componentName] = component
}

func (e *Entity) RemoveComponent(componentName string) {
	_, ok := e.Components[componentName]
	if ok {
		delete(e.Components, componentName)
	}
}


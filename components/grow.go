package components

type GrowComponent struct { 
    Growth int
}

func (c GrowComponent) IsComponent() bool {
	return true
}

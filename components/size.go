package components

type SizeComponent struct {
	Height int
}

func (c SizeComponent) IsComponent() bool {
	return true
}

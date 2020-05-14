package gun

import (
	factory2 "pattern/creational/factory"
)

type maverick struct {
	factory2.Gun
}

func NewMaverick() factory2.IGun {
	return &maverick{
		Gun: factory2.Gun{
			Name:  "Maverick gun",
			Power: 5,
		},
	}
}

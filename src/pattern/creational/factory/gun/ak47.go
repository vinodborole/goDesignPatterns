package gun

import (
	factory2 "pattern/creational/factory"
)

type ak47 struct {
	factory2.Gun
}

func NewAk47() factory2.IGun {
	return &ak47{
		Gun: factory2.Gun{
			Name:  "AK47 gun",
			Power: 4,
		},
	}
}

package factory

type Gun struct {
	Name  string
	Power int
}

func (g *Gun) setName(name string) {
	g.Name = name
}

func (g *Gun) GetName() string {
	return g.Name
}

func (g *Gun) setPower(power int) {
	g.Power = power
}

func (g *Gun) GetPower() int {
	return g.Power
}
/*
*	Concrete product
*/

package factorymethod

type Gun struct {
	name string
	power int
}

func (g *Gun) setName(_name string) {
	g.name = _name
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) setPower(_power int) {
	g.power = _power
}

func (g *Gun) getPower() int {
	return g.power
}
/**
* Concrete flyweight object
 */

package flyweight

type CounterTerroristDress struct {
	color string
}

func (c *CounterTerroristDress) getColor() string {
	return c.color
}

func newCountTerroristDress() *CounterTerroristDress {
	return &CounterTerroristDress{
		color: "green",
	}
}

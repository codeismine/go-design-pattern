/**
* Client code
*/

package flyweight

type game struct {
	terrorists []*Player
	counterTerrorists []*Player
}

func newGame() *game {
	return &game{
		terrorists: make([]*Player, 1),
		counterTerrorists: make([]*Player, 1),
	}
}

func (g *game) addTerrorist(dressType string) {
	player := newPlayer("T", dressType)
	g.terrorists = append(g.terrorists, player)
}

func (g *game) addCounterTerrorist(dressType string) {
	player := newPlayer("CT", dressType)
	g.counterTerrorists = append(g.counterTerrorists, player)
}
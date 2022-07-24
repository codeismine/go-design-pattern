/*
*	Abstract product
*/

package abstractfactory

type IShirt interface {
	setLogo(_logo string)
	setSize(_size int)
	getLogo() string
	getSize() int
}

type Shirt struct {
	logo string
	size int
}

func (s *Shirt) setLogo(_logo string) {
	s.logo = _logo
}

func (s *Shirt) getLogo() string {
	return s.logo
}

func (s *Shirt) setSize(_size int) {
	s.size = _size
}

func (s *Shirt) getSize() int {
	return s.size
}
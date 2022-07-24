/*
*	Abstract product
*/

package abstractfactory

type IShoe interface {
	setLogo(_logo string)
	setSize(_size int)
	getLogo() string
	getSize() int
}

type Shoe struct {
	logo string
	size int
}

func (s *Shoe) setLogo(_logo string) {
	s.logo = _logo
}

func (s *Shoe) getLogo() string {
	return s.logo
}

func (s *Shoe) setSize(_size int) {
	s.size = _size
}

func (s *Shoe) getSize() int {
	return s.size
}
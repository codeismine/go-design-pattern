/*
*	Product interface
 */

package factorymethod

type IGun interface {
	setName(_name string)
	setPower(_power int)
	getName() string
	getPower() int
}

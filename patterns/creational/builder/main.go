/*
*	Client code
 */

package builder

import "fmt"

func BuilderInitializer() {
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()
	printHouse("Normal", normalHouse)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()
	printHouse("Igloo", iglooHouse)
}

func printHouse(houseType string, h House) {
	fmt.Printf("%s House Door Type: %s\n", houseType, h.doorType)
	fmt.Printf("%s House Window Type: %s\n", houseType, h.doorType)
	fmt.Printf("%s House Num Floor: %d\n", houseType, h.floor)
}
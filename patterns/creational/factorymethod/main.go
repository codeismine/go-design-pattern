/*
*	Client code
 */

package factorymethod

import "fmt"

func BuilderInitializer() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g IGun) {
	fmt.Printf("Gun: %s\n", g.getName())
	fmt.Printf("Power: %d\n", g.getPower())
}

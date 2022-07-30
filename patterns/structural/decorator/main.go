/**
* Client code
 */

package decorator

import "fmt"

func DecoratorInitializer() {
	pizza := &VeggieMania{}

	// Add cheese topping
	pizzaWithCheese := &CheeseTopping{
		pizza: pizza,
	}

	// Add tomato topping
	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggieMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
}

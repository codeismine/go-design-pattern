package main

import (
	"fmt"

	"github.com/codeismine/go-design-pattern/patterns/creational/abstractfactory"
	"github.com/codeismine/go-design-pattern/patterns/creational/builder"
	"github.com/codeismine/go-design-pattern/patterns/creational/factorymethod"
)

func main() {
	// Abstract Factory
	printStart("Abstract Factory")
	abstractfactory.AbstractFactoryInitializer()
	printEnd()

	// Builder 
	printStart("Builder")
	builder.BuilderInitializer()
	printEnd()

	// Factory Method
	printStart("Factory Method")
	factorymethod.BuilderInitializer()
	printEnd()
}

func printStart(title string) {
	fmt.Printf("%s:\n", title)
	fmt.Println("--------------------------")
}

func printEnd() {
	fmt.Printf("--------------------------\n\n\n")
}
package main

import (
	"fmt"

	"github.com/codeismine/go-design-pattern/patterns/creational/abstractfactory"
	"github.com/codeismine/go-design-pattern/patterns/creational/builder"
)

func main() {
	// Abstract Factory
	printPattern("Abstract Factory")
	abstractfactory.AbstractFactoryInitializer()
	printEnd()

	// Builder 
	printPattern("Builder")
	builder.BuilderInitializer()
	printEnd()
}

func printPattern(title string) {
	fmt.Printf("%s:\n", title)
	fmt.Println("--------------------------")
}

func printEnd() {
	fmt.Printf("--------------------------\n\n\n")
}
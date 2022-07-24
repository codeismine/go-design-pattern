package main

import (
	"fmt"

	"github.com/codeismine/go-design-pattern/patterns/creational/abstractfactory"
	"github.com/codeismine/go-design-pattern/patterns/creational/builder"
	"github.com/codeismine/go-design-pattern/patterns/creational/factorymethod"
	"github.com/codeismine/go-design-pattern/patterns/creational/prototype"
	// singleton1 "github.com/codeismine/go-design-pattern/patterns/creational/singleton/exampleone"
	// singleton2 "github.com/codeismine/go-design-pattern/patterns/creational/singleton/exampletwo"
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

	// Prototype
	printStart("Prototype")
	prototype.PrototypeInitializer()
	printEnd()

	// Singleton
	// printStart("Singleton Ex1")
	// singleton1.SingletonInitializer()
	// printEnd()

	// printStart("Singleton Ex2")
	// singleton2.SingletonInitializer()
	// printEnd()
}

func printStart(title string) {
	fmt.Printf("%s:\n", title)
	fmt.Println("--------------------------")
}

func printEnd() {
	fmt.Printf("--------------------------\n\n\n")
}
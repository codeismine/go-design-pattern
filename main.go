package main

import (
	"fmt"

	// "github.com/codeismine/go-design-pattern/patterns/creational/abstractfactory"
	// "github.com/codeismine/go-design-pattern/patterns/creational/builder"
	// "github.com/codeismine/go-design-pattern/patterns/creational/factorymethod"
	// "github.com/codeismine/go-design-pattern/patterns/creational/prototype"
	// singleton1 "github.com/codeismine/go-design-pattern/patterns/creational/singleton/exampleone"
	// singleton2 "github.com/codeismine/go-design-pattern/patterns/creational/singleton/exampletwo"
	// "github.com/codeismine/go-design-pattern/patterns/structural/adapter"
	// "github.com/codeismine/go-design-pattern/patterns/structural/bridge"
	// "github.com/codeismine/go-design-pattern/patterns/structural/composite"
	// "github.com/codeismine/go-design-pattern/patterns/structural/decorator"
	// "github.com/codeismine/go-design-pattern/patterns/structural/facade"
	// "github.com/codeismine/go-design-pattern/patterns/structural/flyweight"
	// "github.com/codeismine/go-design-pattern/patterns/structural/proxy"
	"github.com/codeismine/go-design-pattern/patterns/behavioral/chainofresponsibility"
)

func main() {
	// Abstract Factory
	// printStart("Abstract Factory")
	// abstractfactory.AbstractFactoryInitializer()
	// printEnd()

	// Builder 
	// printStart("Builder")
	// builder.BuilderInitializer()
	// printEnd()

	// Factory Method
	// printStart("Factory Method")
	// factorymethod.BuilderInitializer()
	// printEnd()

	// Prototype
	// printStart("Prototype")
	// prototype.PrototypeInitializer()
	// printEnd()

	// Singleton
	// printStart("Singleton Ex1")
	// singleton1.SingletonInitializer()
	// printEnd()

	// printStart("Singleton Ex2")
	// singleton2.SingletonInitializer()
	// printEnd()

	// Adapter
	// printStart("Adapter")
	// adapter.AdapterInitializer()
	// printEnd()

	// Bridge
	// printStart("Bridge")
	// bridge.BridgeInitializer()
	// printEnd()

	// Composite
	// printStart("Composite")
	// composite.CompositeInitializer()
	// printEnd()

	// Decorator
	// printStart("Decorator")
	// decorator.DecoratorInitializer()
	// printEnd()

	// Facade
	// printStart("Facade")
	// facade.FacadeInitializer()
	// printEnd()

	// Flyweight
	// printStart("Flyweight")
	// flyweight.FlyweightInitializer()
	// printEnd()

	// Proxy
	// printStart("Proxy")
	// proxy.ProxyInitializer()
	// printEnd()

	// Chain of Responsibility
	printStart("Chain of Responsibility")
	chainofresponsibility.ChainOfResponsibilityInitializer()
	printEnd()
}

func printStart(title string) {
	fmt.Printf("%s:\n", title)
	fmt.Println("--------------------------")
}

func printEnd() {
	fmt.Printf("--------------------------\n\n\n")
}
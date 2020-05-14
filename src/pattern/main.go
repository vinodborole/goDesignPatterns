package main

import (
	"fmt"
	"pattern/creational/factory"
	"pattern/creational/factory/gunFactory"
	"pattern/creational/singleton"
)

func main() {
	executeCreationalDesignPatterns()
}

func executeCreationalDesignPatterns(){
	//Factory Patter
	factoryPattern()
	//Singleton Pattern
	singletonPattern()
}

func singletonPattern() {
	for i := 0; i < 100; i++ {
		go singleton.GetDatabaseInstance()
	}
}

func factoryPattern() {
	ak47, _ := gunFactory.GetGun("ak47")
	maverick, _ := gunFactory.GetGun("maverick")
	printDetails(ak47)
	printDetails(maverick)
}

func printDetails(g factory.IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}

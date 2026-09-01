package main

import (
	"fmt"
	"slices"
)

const DESCRIPTION string = "This is a package that describes a dog."

var BEST_DOG_BREEDS = []string{"Golden Retriever", "Border Collie"}

type Dog struct {
	name  string
	breed string
}

func newDog(name string, breed string) *Dog {
	return &Dog{name, breed}
}

func (dog *Dog) isBestBreed() bool {
	return slices.Contains(BEST_DOG_BREEDS, dog.breed)
}

func (dog *Dog) isBestDog() {
	if dog.isBestBreed() {
		fmt.Printf("%s is the best dog, congrats!\n", dog.name)
	} else {
		fmt.Printf("%s is not the best dog, sorry.\n", dog.name)
	}
}

func Description() {
	fmt.Println(DESCRIPTION)
}

func main() {
	Description()
	var noah *Dog = newDog("Noah", "Golden Retriever")
	var nina *Dog = newDog("Nina", "Border Collie")
	var pepe *Dog = newDog("Pepe", "Pitbull")
	noah.isBestDog()
	nina.isBestDog()
	pepe.isBestDog()
}

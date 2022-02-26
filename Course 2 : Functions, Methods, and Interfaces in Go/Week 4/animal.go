package main

import (
	"fmt"
	"errors"
)

type Named interface {
	getName() string
}

type Animal interface {
	Named
	Eat()
	Move()
	Speak()
}

// Cow type
type Cow struct {
	name string
}

// Bird type
type Bird struct {
	name string
}

// Snake type
type Snake struct {
	name string
}

func (cow Cow) getName() string {
	return cow.name
}

func (bird Bird) getName() string {
	return bird.name
}

func (snake Snake) getName() string {
	return snake.name
}

// Eat method from Animal interface
func (cow Cow) Eat() {
	fmt.Println("grass")
}

// Eat method from Animal interface
func (bird Bird) Eat() {
	fmt.Println("worms")
}

// Eat method from Animal interface
func (snake Snake) Eat() {
	fmt.Println("mice")
}

// Move method from Animal interface
func (cow Cow) Move() {
	fmt.Println("walk")
}

// Move method from Animal interface
func (bird Bird) Move() {
	fmt.Println("fly")
}

// Move method from Animal interface
func (snake Snake) Move() {
	fmt.Println("slither")
}

// Speak method from Animal interface
func (cow Cow) Speak() {
	fmt.Println("moo")
}

// Speak method from Animal interface
func (bird Bird) Speak() {
	fmt.Println("peep")
}

// Speak method from Animal interface
func (snake Snake) Speak() {
	fmt.Println("hsss")
}

func findAnimal(animals []Animal, name string) (Animal, error) {
	for _, animal := range animals {
		if Named(animal).getName()== name {
			return animal,nil
		}
	}
	return nil, errors.New("Animal with that name not found")
}

func main(){
	var animals []Animal

	for {
		var command_type,name,option string

		fmt.Print(">")
		if _, err := fmt.Scanln(&command_type, &name, &option); err != nil {
			fmt.Println("Error:", err)
			break
		}

		switch command_type {
		case "newanimal":
			switch option {
			case "cow":
				animals = append(animals, Cow{name})
			case "bird":
				animals = append(animals, Bird{name})
			case "snake":
				animals = append(animals, Snake{name})
			}
		case "query":
			var animal Animal
			var err error
			if animal, err = findAnimal(animals, name); err != nil {
				fmt.Println(err)
				break
			}
			switch option {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			}
		default :
			fmt.Println("Wrong Command")
		}

	}
}
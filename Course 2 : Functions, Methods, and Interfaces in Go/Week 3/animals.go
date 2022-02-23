package main

import (
	"fmt"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal *Animal) Eat() string{
	return animal.food
}

func (animal *Animal) Move() string{
	return animal.locomotion
}

func (animal *Animal) Speak() string{
	return animal.noise
}

func main(){

	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	for {
		var animalChoice,infoChoice string

		fmt.Print(">")
		if _, err := fmt.Scanln(&animalChoice, &infoChoice); err != nil {
			fmt.Println("Error:", err)
			break
		}

		var animal Animal
		if animalChoice == "cow" {
			animal = cow
		} else if animalChoice == "bird" {
			animal = bird
		} else if animalChoice == "snake" {
			animal = snake
		}

		if infoChoice == "eat" {
			fmt.Println(animal.Eat())
		} else if infoChoice == "move" {
			fmt.Println(animal.Move())
		} else if infoChoice == "speak" {
			fmt.Println(animal.Speak())
		}


	}
}
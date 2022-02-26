package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (c Cow) Eat() {
	fmt.Println("grass")
}
func (c Cow) Move() {
	fmt.Println("walk")
}
func (c Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct{}

func (c Bird) Eat() {
	fmt.Println("worms")
}
func (c Bird) Move() {
	fmt.Println("fly")
}
func (c Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct{}

func (c Snake) Eat() {
	fmt.Println("mice")
}
func (c Snake) Move() {
	fmt.Println("slither")
}
func (c Snake) Speak() {
	fmt.Println("hsss")
}

func main() {

	animals := make(map[string]Animal)

	re := regexp.MustCompile("\\s+")
	reader := bufio.NewReader(os.Stdin)

	//   fmt.Println("(x to quit)")
	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		params := re.Split(text, -1)
		if len(params) != 3 {
			fmt.Println("wrong input, it forms newanimal name animal or query name info")
			continue
		}

		command := params[0]
		name := params[1]
		what := params[2]

		switch command {
		case "newanimal":
			switch what {
			case "cow":
				animals[name] = Cow{}
			case "bird":
				animals[name] = Bird{}
			case "snake":
				animals[name] = Snake{}
			default:
				fmt.Printf("wrong animal: %s, should be one of cow, bird and snake\n", what)
			}
		case "query":
			if animal, ok := animals[name]; ok {
				switch what {
				case "eat":
					animal.Eat()
				case "move":
					animal.Move()
				case "speak":
					animal.Speak()
				default:
					fmt.Printf("no such info exists: %s\n", what)
				}
			} else {
				fmt.Println("no animla found for the name: %s\n", name)
			}
		default:
			fmt.Println("unsupported command %s\n", command)
		}
	}
}

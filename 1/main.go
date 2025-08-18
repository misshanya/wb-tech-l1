package main

import "fmt"

type Human struct {
	Name string
	Age  int
	Sex  string
}

func (h Human) Greet() {
	fmt.Println("Hi! I'm", h.Name)
}

type Action struct {
	Human
}

func main() {
	human := Human{
		Name: "Pavel",
		Age:  20,
		Sex:  "male",
	}

	human.Greet()

	action := Action{
		Human{
			Name: "Pavel Pavlov",
			Age:  20,
			Sex:  "male",
		},
	}

	action.Greet()
}

package main

import "fmt"

// Greeter defines someone
type Greeter struct {
	Format string
}

// Greet greets a human (or robot)
func (g Greeter) Greet(name string) {
	fmt.Printf(g.Format, name)
}

func run1() {
	g := Greeter{
		Format: "Sup %s\n",
	}

	g.Greet("sergi")

	Greeter.Greet(g, "sergi")
}

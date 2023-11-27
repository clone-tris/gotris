package main

import (
	"fmt"
	"gotris/screens/game/components"
)

func main() {
	s := components.Square{Row: 1, Column: 1}
	fmt.Printf("%+v\n", s)
}

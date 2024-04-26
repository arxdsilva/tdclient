package main

import (
	"fmt"

	"github.com/arxdsilva/tdclient"
)

func main() {
	client := tdclient.New()
	world, err := client.GetWorld("Premia")
	if err != nil {
		// handle error
	}

	worlds, err := client.GetWorlds()
	if err != nil {
		// handle error
	}

	fmt.Println("world: ", world)
	fmt.Println("worlds: ", worlds)
}

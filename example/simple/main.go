package main

import (
	"context"
	"fmt"

	"github.com/arxdsilva/tdclient"
)

func main() {
	client := tdclient.New()
	world, err := client.GetWorld(context.Background(), "Premia")
	if err != nil {
		// handle error
	}

	worlds, err := client.GetWorlds(context.Background())
	if err != nil {
		// handle error
	}

	fmt.Println("world: ", world)
	fmt.Println("worlds: ", worlds)
}

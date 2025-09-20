package main

import (
	"fmt"

	"github.com/jmjp/hltv"
)

func main() {
	player, err := hltv.FetchPlayerById(2023)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Player: %s\n", player.Name)
}

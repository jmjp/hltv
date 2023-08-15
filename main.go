package main

import (
	"fmt"

	"github.com/jmjp/hltv/hltv"
)

func main() {
	_, err := hltv.FetchMatches()
	if err != nil {
		fmt.Println(err)
	}
}

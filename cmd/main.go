package main

import (
	"fmt"

	"github.com/jmjp/hltv"
)

func main() {
	matches, err := hltv.FetchMatches()
	if err != nil {
		panic(err)
	}
	for _, match := range matches {
		teamA := "TBO"
		teamB := "TBO"
		time := match.Date.String()
		if match.TeamA != nil {
			teamA = match.TeamA.Name
		}
		if match.TeamB != nil {
			teamB = match.TeamB.Name
		}
		if match.Live {
			time = "LIVE"
		}
		fmt.Printf("[%s] %s vs %s\n", time, teamA, teamB)
	}
}

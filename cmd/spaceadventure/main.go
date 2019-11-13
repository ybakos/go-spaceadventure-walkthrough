package main

import "github.com/ybakos/spaceadventure_walkthrough/internal/spaceadventure"

func main() {
	spaceadventure.Start(
		spaceadventure.PlanetarySystem{
			Name: "Solar System", Planets: []spaceadventure.Planet{
				spaceadventure.Planet{"Tatooine", "Desert planet"},
			},
		},
	)
}

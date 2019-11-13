package spaceadventure

import "fmt"

func Start(planetarySystem PlanetarySystem) {
	printWelcome(planetarySystem)
	printGreeting(responseToPrompt("What is your name?"))
	fmt.Println("Let's go on an adventure!")
	travel(promptForRandomOrSpecificDestination("Shall I randomly choose a planet for you to visit? (Y or N)"))
}

func printWelcome(planetarySystem PlanetarySystem) {
	fmt.Printf("Welcome to the %s!\n", planetarySystem.Name)
	fmt.Printf("There are %d planets to explore.\n", planetarySystem.NumberOfPlanets())
}

func printGreeting(name string) {
	fmt.Printf("Nice to meet you, %s. My name is Eliza, I'm an old friend of Siri.\n", name)
}

func responseToPrompt(prompt string) (response string) {
	fmt.Println(prompt)
	fmt.Scan(&response)
	return
}

func promptForRandomOrSpecificDestination(prompt string) bool {
	var choice string
	for choice != "Y" && choice != "N" {
		choice = responseToPrompt(prompt)
		if choice == "Y" {
			return true
		} else if choice == "N" {
			return false
		} else {
			fmt.Println("Sorry, I didn't get that.")
		}
	}
	return false
}

func travel(randomDestination bool) {
	if randomDestination {
		travelToRandomPlanet()
	} else {
		travelToPlanet(responseToPrompt("Name the planet you would like to visit."))
	}
}

func travelToRandomPlanet() {
	fmt.Println("Traveling to Jupiter...")
	fmt.Println("Arrived at Jupiter. The large red spot appears ominous.")
}

func travelToPlanet(planetName string) {
	fmt.Printf("Traveling to %s...\n", planetName)
	fmt.Println("Arrived at Neptune. A very cold planet, furthest from the sun.")
}

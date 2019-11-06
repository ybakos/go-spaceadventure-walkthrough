package spaceadventure

import "fmt"

func PrintWelcome() {
	fmt.Println("Welcome to the Solar System!")
	fmt.Println("There are 8 planets to explore.")
}

func ResponseToPrompt(prompt string) string {
	var response string
	fmt.Println(prompt)
	fmt.Scan(&response)
	return response
}

func PrintGreeting(name string) {
	fmt.Printf("Nice to meet you, %s. My name is Eliza, I'm an old friend of Siri.\n", name)
}

func Travel() {
	var choice string
	for choice != "Y" && choice != "N" {
		choice = ResponseToPrompt("Shall I randomly choose a planet for you to visit? (Y or N)")
		if choice == "Y" {
			TravelToRandomPlanet()
		} else if choice == "N" {
			TravelToPlanet(ResponseToPrompt("Name the planet you would like to visit."))
		} else {
			fmt.Println("Sorry, I didn't get that.")
		}
	}
}

func TravelToRandomPlanet() {
	fmt.Println("Traveling to Jupiter...")
	fmt.Println("Arrived at Jupiter. The large red spot appears ominous.")
}

func TravelToPlanet(planetName string) {
	fmt.Printf("Traveling to %s...\n", planetName)
	fmt.Println("Arrived at Neptune. A very cold planet, furthest from the sun.")
}

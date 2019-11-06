package main

import "fmt"
import "github.com/ybakos/spaceadventure_walkthrough/internal/spaceadventure"

func main() {
	spaceadventure.PrintWelcome()
	spaceadventure.PrintGreeting(spaceadventure.ResponseToPrompt("What is your name?"))
	fmt.Println("Let's go on an adventure!")
	spaceadventure.Travel()
}

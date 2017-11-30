package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var responses = []string{
		"People say I look like both my mother and father.",
		"Father was a teacher.",
		"I was my father’s favourite.",
		"I’m looking forward to the weekend.",
		"My grandfather was French!",
	}

	for i := 0; i < len(responses); i++ {
		fmt.Println(responses[i])
		fmt.Println(ElizaResponse(responses[i]))
	}
}

// ElizaResponse takes a single string of input and returns a single string of output
func ElizaResponse(input string) string {
	var responses = []string{
		"I'm not sure what you're trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
	}

	return randomChoice(responses)
}

// returns a random string from a list of strings
func randomChoice(input []string) string {
	randIndex := rand.Intn(len(input))
	return input[randIndex]
}

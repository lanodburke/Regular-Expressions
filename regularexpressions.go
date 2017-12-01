package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
)

func main() {
	var responses = []string{
		"People say I look like both my mother and father.",
		"Father was a teacher.",
		"I was my father’s favourite.",
		"I’m looking forward to the weekend.",
		"My grandfather was French!",
		"I am happy.",
		"I am not happy with your responses.",
		"I am not sure that you understand the effect that your questions are having on me.",
		"I am supposed to just take what you’re saying at face value?",
	}

	for i := 0; i < len(responses); i++ {
		fmt.Println(responses[i])
		fmt.Println(ElizaResponse(responses[i]))
	}
}

var reflection = map[string]string{
	"am":     "are",
	"was":    "were",
	"i":      "you",
	"i'd":    "you would",
	"i've":   "you have",
	"i'll":   "you will",
	"my":     "your",
	"are":    "am",
	"you've": "I have",
	"you'll": "I will",
	"your":   "my",
	"yours":  "mine",
	"you":    "me",
	"me":     "you",
}

// ElizaResponse takes a single string of input and returns a single string of output
func ElizaResponse(input string) string {
	var responses = []string{
		"I'm not sure what you're trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
	}

	regex, _ := regexp.Compile("father")

	if regex.MatchString(input) {
		return "Why don’t you tell me more about your father?"
	}

	re := regexp.MustCompile(`(?i)i(?:'|\sa)?m (.*)`)
	matched := re.FindStringSubmatch(input)

	if len(matched) > 0 {
		var res string
		if len(matched) > 1 {
			res = reflect(matched[1])
		}

		response := randomChoice(responses)
		if strings.Contains(response, "%s") {
			response = fmt.Sprintf(response, res)
		}
		return response
	}

	return randomChoice(responses)
}

func reflect(input string) string {
	words := strings.Split(input, " ")
	for i, word := range words {
		if reflected, ok := reflection[word]; ok {
			words[i] = reflected
		}
	}
	return strings.Join(words, " ")
}

// returns a random string from a list of strings
func randomChoice(input []string) string {
	randIndex := rand.Intn(len(input))
	return input[randIndex]
}

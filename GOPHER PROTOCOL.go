package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Available commands: (upper, lower, cap, snake, title, reverse, exit): ")
		fmt.Print("Enter word To Transform: ")
		word, _ := reader.ReadString('\n')
		word = strings.TrimSpace(word)

		if word == "" {
			continue
		}

		text := strings.Fields(word)

		command := text[0]
		textslice := text[1:]

		fword := strings.Join(textslice, " ")

		if len(fword) <= 1 {
			fmt.Println("No text provided. Usage: <command> <text>")
			continue
		}

	


		switch strings.ToLower(command) {
		case "upper":
			fmt.Println("Result: ", strings.ToUpper(fword))
		case "lower":
			fmt.Println("Result: ", strings.ToLower(fword))
		case "cap":
			words := strings.ToLower(fword)
			fmt.Println("Result: ", (cap(words)))
		case "snake":
			w := strings.ToLower(fword)
			s := strings.ReplaceAll(w, " ", "_")
			fmt.Println("Result: ", s)
		case "exit":
			fmt.Println("Shutting down String Transformer. Goodbye.")
			return
		case "reverse":
			fmt.Println("Result: ", rev(fword))
		case "title":
			fmt.Println(title(fword))
		default:
			fmt.Println("Unknown command")
			fmt.Println("Valid commands: upper, lower, cap,title, snake, reverse, exit")
			fmt.Println("Try Again")
		}
	}
}

func upper(s string) string {
	return strings.ToUpper(s)
}

func lower(s string) string {
	return strings.ToLower(s)
}

func cap(s string) string {
	word := strings.Fields(s)

	for i, words := range word {
		word[i] = strings.ToUpper(string(words[0])) + strings.ToLower(string(words[1:]))
	}
	return strings.Join(word, " ")
}

func snake_case(s string) string {
	w := strings.ToLower(s)
	word := strings.ReplaceAll(w, " ", "_")
	return string(word)

}

func rev(s string) string {
	w := []rune(s)
	for i, j := 0, len(w)-1; i < j; i, j = i+1, j-1 {
		w[i], w[j] = w[j], w[i]

	}
	return string(w)
}

func title(text string)  string {
	w := map[string]bool {
		"a": true, "the": true, "and": true, "but": true, "of": true, "nor": true,
		"for": true, "it": true,  "or": true,  "as": true, "is": true, "by": true,
		"an": true, "on": true, "at": true, "to": true, "in": true, "up": true, 
	}
	words:= strings.Fields(strings.ToLower(text))
	for i, word := range words {
		if i == 0 || !w[word] {
			words[i] = strings.ToUpper(string(word[0])) + word[1:]
		}
	}
	return strings.Join(words, " ")
}

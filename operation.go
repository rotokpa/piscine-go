package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Hey! Enter word: ")
		word, _ := reader.ReadString('\n')
		word = strings.TrimSpace(word)

		if word == "exit" {
			fmt.Println("Goodbye")
			break
		}

		isvalid := false
		for _, w := range word {
			if unicode.IsLetter(w) || unicode.IsDigit(w) {
				isvalid = false
				break
			}
		}
		if isvalid {
			fmt.Println("only latters are allowed")
			continue
		}
	

		var operation string
		fmt.Print("Enter operation: (reverse, lastchar, capitalize, deleteIndex): ")
		fmt.Scan(&operation)

		switch operation {
		case "reverse":
			w := strings.Fields(word) 
			for i, j := 0, len(w)-1; i < j; i, j = i+1, j-1 {
				w[i], w[j] = w[j], w[i]
				
				res := strings.Join(w, " ")
				fmt.Println("Result: ", res)
				break
			}

		case "lastchar":
			if len(word) > 0 {
				fmt.Println("Result: ", string(word[len(word)-1]))
			}else {
				fmt.Println("word is empty")
			}

		case "capitalize":
			fmt.Println("Result: ", strings.ToUpper(word))

		case "deleteIndex":
			var index int 
			fmt.Print("Enter Index to be deleted : ")
			fmt.Scan(&index)

			if index < 0 || index >= len(word) {
				fmt.Println("Invalid index is out of range: ", len(word[:0]), "and", len(word)-1)
				continue
			}
			result := word[:index] + word[index+1:] 
			fmt.Println("Result: ", result)
		default:
			fmt.Println("Invalid Operation")
		}
	}
}
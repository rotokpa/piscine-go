package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	for {
		var convertion string
		fmt.Print("Enter string to convert or ('exit' to quit program): ")
		fmt.Scan(&convertion)

		if convertion == "exit" {
			fmt.Println("exiting progam...")
			return
		}

		var command string
		fmt.Println("Choose coverter(hex, bin, dec): ")
		fmt.Scanln(&command)

		switch command {
		case "hex":
			val, err := hexTodec(convertion)
			if err != nil {
				fmt.Println("Invalid Hex")
			} else {
				fmt.Println("Decimal: ", val)
			}
		case "bin":
			val, err := binTodec(convertion)
			if err != nil {
				fmt.Println("Invalid bin")
			} else {
				fmt.Println("decimal: ", val)
			}
		case "dec":
			val, err := strconv.ParseInt(convertion, 10, 64)
			if err != nil {
				fmt.Println("Invalid dec")
			} else {
				fmt.Println(strings.ToUpper(dec1(val)))
				fmt.Println(dec2(val))
			}
		}
	}
}

func hexTodec(hexStr string) (int64, error) {
	return strconv.ParseInt(hexStr, 16, 64)
}

func binTodec(binStr string) (int64, error) {
	return strconv.ParseInt(binStr, 2, 64)
}

func dec1(val int64) string {
	return strconv.FormatInt(val, 16)
}

func dec2(val int64) string {
	return strconv.FormatInt(val, 2)
}

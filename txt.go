package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var number = [7]int{1, 2, 3, 4, 5, 6, 7}

var scanner = bufio.NewScanner(os.Stdin)
var input string

func main() {

	fmt.Printf("Enter your Text: ")
	scanner.Scan()
	input = string(scanner.Text())

	transformation()

}

func transformation() {

	for true {

		fmt.Println("What would you like to do?")
		fmt.Println("(1) Backwards", "(2) Random", "(3) Everything great", "(4) Everything small")
		fmt.Println("(5) First letter capitalized", "(6) New Text", "(7) Exit")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case number[0]:
			fmt.Println("Answer: ", backwards(input))

		case number[1]:
			fmt.Println("Answer: ", random(input))

		case number[2]:
			output := strings.ToUpper(input)
			fmt.Println("Answer: " + output)

		case number[3]:
			output := strings.ToLower(input)
			fmt.Println("Answer: " + output)

		case number[4]:
			output := strings.Title(input)
			fmt.Println("Answer: " + output)

		case number[5]:
			main()

		case number[6]:
			os.Exit(0)

		}

	}

}

func random(s string) string {

	rand.Seed(time.Now().UnixNano())

	words := strings.Fields(s)
	for i := len(words) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")

}

func backwards(s string) string {

	words := strings.SplitAfter(s, "")
	newString := make([]string, len(words))

	for k, h := range words {
		newString[(len(words) - k - 1)] = h
	}

	return strings.Join(newString, "")

}

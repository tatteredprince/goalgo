package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner, input := bufio.NewScanner(os.Stdin), ""
	for scanner.Scan() {
		input += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("read error: %e", err)
		os.Exit(1)
	}

	parentheses := []byte(input)

	openParenthesesCount := 0

	for _, s := range parentheses {
		if s == '(' {
			openParenthesesCount++
		} else if openParenthesesCount != 0 && s == ')' {
			openParenthesesCount--
		} else {
			fmt.Println(0)
			os.Exit(0)
		}
	}

	if openParenthesesCount == 0 {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}

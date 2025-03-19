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

	openParenthesCount := 0

	for _, s := range parentheses {
		if s == '(' {
			openParenthesCount++
		} else if openParenthesCount != 0 && s == ')' {
			openParenthesCount--
		} else {
			fmt.Println(0)
			os.Exit(0)
		}
	}

	if openParenthesCount == 0 {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}

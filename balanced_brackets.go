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

	brackets := []byte(input)

	stack := make([]byte, 0, len(brackets))

	for _, s := range brackets {
		iprev := len(stack) - 1

		if s == '(' || s == '{' || s == '[' {
			stack = append(stack, s)
		} else if len(stack) != 0 &&
			(stack[iprev] == '(' && s == ')' ||
				stack[iprev] == '{' && s == '}' ||
				stack[iprev] == '[' && s == ']') {
			stack = stack[:iprev]
		} else {
			fmt.Println(0)
			os.Exit(0)
		}
	}

	if len(stack) == 0 {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}

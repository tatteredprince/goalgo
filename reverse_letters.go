package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	str := ""
	for scanner.Scan() {
		str += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("read error: %e", err)
		os.Exit(1)
	}

	reversed, bytes, stack := "", []byte(str), []byte{}

	cmpByte := func(ch byte) bool { return ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z' }

	for i := 0; i < len(bytes); i++ {
		for ; i < len(bytes) && cmpByte(bytes[i]); i++ {
			stack = append(stack, bytes[i])
		}

		if len(stack) != 0 {
			for j := len(stack) - 1; j >= 0; j-- {
				reversed += string(stack[j])
			}
			stack = []byte{}
			i--
			continue
		}

		reversed += string(bytes[i])
	}

	fmt.Println(reversed)
}

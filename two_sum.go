package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner, input := bufio.NewScanner(os.Stdin), ""
	for scanner.Scan() {
		input += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("read error: %e\n", err)
		os.Exit(1)
	}

	strs := strings.Split(input, " ")

	target, err := strconv.Atoi(strs[len(strs)-1])
	if err != nil {
		fmt.Printf("conversion error: %e\n", err)
		os.Exit(1)
	}

	ints := make([]int, len(strs)-1)
	for i := range len(strs) - 1 {
		ints[i], err = strconv.Atoi(strs[i])
		if err != nil {
			fmt.Printf("conversion error at index %d: %e", i, err)
			os.Exit(1)
		}
	}

	set := map[int]struct{}{}

	for _, v := range ints {
		complement := target - v

		if _, ok := set[complement]; ok == true {
			fmt.Println(1)
			os.Exit(0)
		}

		set[v] = struct{}{}
	}

	fmt.Println(0)
}

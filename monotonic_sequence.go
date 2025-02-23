package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := io.ReadAll(os.Stdin)

	seq := []int{}
	strs := strings.Split(string(bytes), " ")
	for _, s := range strs {
		i, _ := strconv.Atoi(s)
		seq = append(seq, i)
	}

	if len(seq) <= 2 {
		fmt.Println(1)
		os.Exit(0)
	}

	idx, isAscending := 1, true

	for ; idx < len(seq); idx++ {
		if seq[idx] != seq[idx-1] {
			if seq[idx] < seq[idx-1] {
				isAscending = false
			}
			break
		}
	}

	for ; idx < len(seq); idx++ {
		if isAscending && seq[idx] < seq[idx-1] || !isAscending && seq[idx] > seq[idx-1] {
			fmt.Println(0)
			os.Exit(0)
		}
	}

	fmt.Println(1)
}

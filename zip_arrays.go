package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

func zip(arrays ...[]int) [][]int {

	zippedArrsLen := (1<<bits.UintSize)/2 - 1
	for _, arr := range arrays {
		if zippedArrsLen > len(arr) {
			zippedArrsLen = len(arr)
		}
	}

	zippedArrs := make([][]int, 0, zippedArrsLen)

	arrLen := len(arrays)

	for i := 0; i < zippedArrsLen; i++ {
		arr := make([]int, 0, arrLen)

		for _, a := range arrays {
			arr = append(arr, a[i])
		}

		zippedArrs = append(zippedArrs, arr)
	}

	return zippedArrs
}

func zip2(arrays ...[]int) [][]int {

	zippedArrsLen := (1<<bits.UintSize)/2 - 1
	for _, arr := range arrays {
		if zippedArrsLen > len(arr) {
			zippedArrsLen = len(arr)
		}
	}

	zippedArrs := make([][]int, zippedArrsLen)

	arrLen := len(arrays)

	for i := 0; i < zippedArrsLen; i++ {
		zippedArrs[i] = make([]int, arrLen)

		for j := 0; j < arrLen; j++ {
			zippedArrs[i][j] = arrays[j][i]
		}
	}

	return zippedArrs
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	inputStr := ""
	for scanner.Scan() {
		inputStr += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("read error: %e", err)
		os.Exit(1)
	}

	arrays := [][]int{}

	for i := 0; i < len(inputStr); i++ {

		if inputStr[i] == '[' {
			for j := i + 1; j < len(inputStr); j++ {

				if inputStr[j] == ']' {
					substr := inputStr[i+1 : j]

					if substr != "" {
						arrStrs := strings.Split(substr, " ")
						arrInts := make([]int, 0, len(arrStrs))

						for _, s := range arrStrs {
							d, err := strconv.Atoi(s)
							if err != nil {
								fmt.Printf("conversion error: %e", err)
								os.Exit(1)
							}

							arrInts = append(arrInts, d)
						}

						arrays = append(arrays, arrInts)
					} else {
						arrays = append(arrays, []int{})
					}

					i = j
					break
				}
			}
		}
	}

	outputStr := fmt.Sprint(zip2(arrays...))

	if outputStr[0:2] == "[[" {
		fmt.Println(outputStr[1 : len(outputStr)-1])
	} else {
		fmt.Println(outputStr)
	}
}

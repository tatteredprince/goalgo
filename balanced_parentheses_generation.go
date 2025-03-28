package main

import (
	"fmt"
)

func generateParentheses(n, open, close int, str string) []string {
	seq := make([]string, 0, n)

	if open == n && close == n {
		seq = append(seq, str)
		return seq
	}

	if open < n {
		seq = append(seq, generateParentheses(n, open+1, close, str+"(")...)
	}

	if close < open {
		seq = append(seq, generateParentheses(n, open, close+1, str+")")...)
	}

	return seq
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	if n != 0 {
		for _, str := range generateParentheses(n, 0, 0, "") {
			fmt.Println(str)
		}
	}
}

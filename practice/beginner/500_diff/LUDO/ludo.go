// SPDX-FileCopyrightText: 2025 Vladimir Rusinov
// SPDX-License-Identifier: Apache-2.0

// Package main implements Chef Plays Ludo
// https://www.codechef.com/practice/course/basic-programming-concepts/DIFF500/problems/LUDO
package main

import "fmt"

func ludo(x int) bool {
	return x == 6
}

func main() {
	var t int
	_, err := fmt.Scan(&t)
	if err != nil {
		panic(err)
	}
	for i := 0; i < t; i++ {
		var x int
		_, err := fmt.Scan(&x)
		if err != nil {
			panic(err)
		}
		if ludo(x) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

// SPDX-FileCopyrightText: 2025 Vladimir Rusinov
// SPDX-License-Identifier: Apache-2.0

// Package main implements Donation Drive
// https://www.codechef.com/practice/course/basic-programming-concepts/DIFF500/problems/DONDRIVE
package main

import "fmt"

func main() {
	var t int

	_, err := fmt.Scanf("%d", &t)
	if err != nil {
		panic(err)
	}

	var n int
	var x int
	for i := 0; i < t; i++ {
		_, err := fmt.Scanf("%d %d", &n, &x)
		if err != nil {
			panic(err)
		}
		fmt.Println(n - x)
	}
}

// SPDX-FileCopyrightText: 2025 Vladimir Rusinov
// SPDX-License-Identifier: Apache-2.0

// How many unattempted problems
// https://www.codechef.com/practice/course/basic-programming-concepts/DIFF500/problems/PRACLIST

package main

import "fmt"

func how_many_unattempted(x int, y int) int {
	return x - y
}

func main() {
	var x int
	var y int
	_, err := fmt.Scanf("%d %d", &x, &y)
	if err != nil {
		panic(err)
	}
	fmt.Println(how_many_unattempted(x, y))
}

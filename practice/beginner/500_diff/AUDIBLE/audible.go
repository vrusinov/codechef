// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: Copyright 2025 Vladimir Rusinov

// package main implements "Audible Range"
// https://www.codechef.com/practice/course/basic-programming-concepts/DIFF500/problems/AUDIBLE
package main

import "fmt"

func main() {
	// Read the number of test cases
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

		if x >= 67 && x <= 45000 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

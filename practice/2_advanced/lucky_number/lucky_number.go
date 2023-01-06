// SPDX-FileCopyrightText: 2022 Google Inc
// SPDX-FileCopyrightText: 2022 Vladimir Rusinov <vrusinov@google.com>
//
// SPDX-License-Identifier: Apache-2.0

// Strategy:
// First attempt to remove all numbers which are multiples of both a and b (to starve the other player). Then remove one at a time.

package main
import (
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"os"
)

func luckyGame(a, b int, numbers []int) string {
	// Don't actually need to remove numbers, just count divisible by both, by a and by b.
	divAB := 0
	divA := 0
	divB := 0
	for _, n := range numbers {
		if ((n % a == 0) && (n % b == 0)) {
			divAB++
		} else if (n % a == 0) {
			divA++
		} else if (n % b == 0) {
			divB++
		}
	}
	// Bob will have div_ab + div_a moves.
	if (divAB + divA > divB) {
		return "BOB"
	}
	return "ALICE"
}

func readInt(reader *bufio.Reader) (int, error) {
	s, err := readLine(reader)
	if err != nil {
	  return 0, err
	}
	return strconv.Atoi(s)
}

func readLine(reader *bufio.Reader) (string, error) {
    str, _, err := reader.ReadLine()
    if err != nil {
        return "", err
    }

    return strings.TrimRight(string(str), "\r\n"), nil
}

func readIntSlice(reader *bufio.Reader) ([]int, error) {
	var f []int
	s, err := readLine(reader)
	if err != nil {
	  return nil, err
	}
	parts := strings.Split(s, " ")
	for _, val := range parts {
	  fN, err := strconv.Atoi(val)
	  if err != nil {
		return nil, err
	  }
	  f = append(f, fN)
	}
	return f, nil
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 2 * 1024 * 1024)

	// The first line of the input contains a single integer TT denoting the number of test cases.
	t, err := readInt(reader)
	if err != nil {
		panic(err)
	}

	// The description of TT test cases follows.
	for i := 0; i < t; i++ {
	  // The first line of each test case contains three space-separated integers NN, aa and bb.
	  s, err := readIntSlice(reader)
	  if err != nil {
		panic(err)
	  }
	  // not used:
	  // n := s[0]
	  a := s[1]
	  b := s[2]
	  numbers, err := readIntSlice(reader)
	  if err != nil {
		panic(err)
	  }
	  fmt.Println(luckyGame(a, b, numbers))
	}
}

// SPDX-FileCopyrightText: 2022 Google Inc
// SPDX-FileCopyrightText: 2022 Vladimir Rusinov <vrusinov@google.com>
//
// SPDX-License-Identifier: Apache-2.0

// Package main implements Shopping Change problem
// https://www.codechef.com/problems/SHOPCHANGE
package main

import (
  "bufio"
  "os"
  "strconv"
  "strings"
  "fmt"
)

const (
	totalMoney = 100
)

func calculateChange(x int) int {
	return totalMoney - x
}

func readLine(reader *bufio.Reader) (string, error) {
    str, _, err := reader.ReadLine()
    if err != nil {
        return "", err
    }

    return strings.TrimRight(string(str), "\r\n"), nil
}

func readInt(reader *bufio.Reader) (int, error) {
  s, err := readLine(reader)
  if err != nil {
    return 0, err
  }
  return strconv.Atoi(s)
}

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

  t, err := readInt(reader)
  if err != nil {
      panic(err)
  }

  for i := 0; i < t; i++ {
    x, err := readInt(reader)
    if err != nil {
      panic(err)
    }
    fmt.Println(calculateChange(x))
  }
}

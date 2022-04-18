// Package main implements Gasoline Introduction problem
// https://www.codechef.com/problems/BEGGASOL
package main

import (
  "bufio"
  "os"
  "strconv"
  "strings"
  "fmt"
)

func gasolineIntroduction(n int, f []int) int {
  carFuel := 0
  distance := 0
  for _, fI := range f {
    // Do one "step"
    carFuel = carFuel + fI - 1
    if carFuel < 0 {
      // Ran out of fuel
      return distance
    }
    distance += fI
  }
  // Did not ran out of fuel, lapping other cars
  return distance
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
  reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
  //writer := bufio.NewWriterSize(os.Stdout, 1024 * 1024)

  t, err := readInt(reader)
  if err != nil {
      panic(err)
  }

  for i := 0; i < t; i++ {
    n, err := readInt(reader)
    if err != nil {
      panic(err)
    }
    f, err := readIntSlice(reader)
    if err != nil {
      panic(err)
    }
    fmt.Println(gasolineIntroduction(n, f))
  }
}

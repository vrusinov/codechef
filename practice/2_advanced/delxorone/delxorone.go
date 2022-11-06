package main
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func delxorone(n int, numbers []int) int {
	// The only way array satisfies condition is when there are only integers with all bits the same except maybe the last one.
	// So we need to calculate what are the most common bits ignoring the last one.
	// That would be two integers we won't remove.

	freqMap := make(map[int]int)
	max := 0
	for _, i := range numbers {
		// Go's default behavior is to return the "zero value" for the value type when a looked-up key is missing.
		temp := freqMap[i / 2] + 1
		freqMap[i / 2] = temp
		if temp > max {
			max = temp
		}
	}
	return n - max
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

    return strings.TrimRight(string(str), "\r\n "), nil
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

	// The first line of the input contains a single integer T denoting the number of test cases. The description of TT test cases follows.
	t, err := readInt(reader)
	if err != nil {
		panic(err)
	}

	// The description of T test cases follows.
	for i := 0; i < t; i++ {
	  // The first line of each test case contains a single integer N.
	  n, err := readInt(reader)
	  if err != nil {
		panic(err)
	  }
	  // The second line contains N space-separated integers
	  numbers, err := readIntSlice(reader)
	  if err != nil {
		panic(err)
	  }
	  fmt.Println(delxorone(n, numbers))
	}
}
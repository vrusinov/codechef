// https://www.codechef.com/submit/MAGICJAR

package main
import (
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"os"
)

func magicjar(n int, numbers []int) int {
	var sum int;
	for _, a := range numbers {
		sum += a;
	}
	// Worst case: everyone has A-1 jars.
	// To avoid we'd need this:
	return sum - n + 1;
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

	// The first line of the input contains a single integer TT denoting the number of test cases. The description of TT test cases follows.
	t, err := readInt(reader)
	if err != nil {
		panic(err)
	}

	// The description of TT test cases follows.
	for i := 0; i < t; i++ {
	  // The first line of each test case contains a single integer NN.
	  n, err := readInt(reader)
	  if err != nil {
		panic(err)
	  }
	  // The second line contains NN space-separated integers
	  numbers, err := readIntSlice(reader)
	  if err != nil {
		panic(err)
	  }
	  fmt.Println(magicjar(n, numbers))
	}
}
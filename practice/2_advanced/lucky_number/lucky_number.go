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

// removeNumbers attempts to remove numbers divisible by both mine and others.
func removeNumbersGreedy(mine, others int, numbers []int) (int, []int) {
	removed := 0
	result := make([]int, 0, len(numbers))
	for _, number := range numbers {
		if ((number % mine == 0) && (number % others == 0)) {
			removed += 1
			continue
		}
		result = append(result, number)
	}

	return removed, result
}

// removeNumbers attempts to remove numbers divisible by both mine and others.
func removeOneNumber(mine int, numbers []int) (int, []int) {
	for i, a := range numbers {
		if (a % mine == 0) {
			return 1, append(numbers[:i], numbers[i+1:]...)
		}
	}

	return 0, numbers
}

func luckyGame(a, b int, numbers []int) string {
	// fmt.Printf("Start array: %v\n", numbers)
	// Bob plays first
	removed := 0
	removed, numbers = removeNumbersGreedy(a, b, numbers)
	// fmt.Printf("Bob removed %v, array is %v\n", removed, numbers)
	if (removed == 0) {
		// Nothing to remove greedily, try to remove one.
		removed, numbers = removeOneNumber(a, numbers)
		if (removed == 0) {
			// Failed to remove any. Alice wins.
			return "ALICE"
		}
	}
	// fmt.Printf("Bob removed %v, array is %v\n", removed, numbers)
	// Now Alice plays:
	removed, numbers = removeNumbersGreedy(b, a, numbers)
	if (removed == 0) {
		// Nothing to remove greedily, try to remove one.
		removed, numbers = removeOneNumber(b, numbers)
		if (removed == 0) {
			// Failed to remove any. Bob wins.
			return "BOB"
		}
	}
	// fmt.Printf("Alice removed %v, array is %v\n", removed, numbers)

	// Now each removes one until somebody wins.
	for {
		// Starting with Bob since Alice played.
		removed, numbers = removeOneNumber(a, numbers)
		if (removed == 0) {
			// Failed to remove any. Alice wins.
			return "ALICE"
		}
		removed, numbers = removeOneNumber(b, numbers)
		if (removed == 0) {
			// Failed to remove any. Bob wins.
			return "BOB"
		}
	}
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
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
	//writer := bufio.NewWriterSize(os.Stdout, 1024 * 1024)

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
//	Print duplicate lines from stdin along with their counts.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	//	Read lines from stdin.

	for input.Scan() {
		counts[input.Text()]++
	}

	//	Ignores potential errors from input.Err().

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

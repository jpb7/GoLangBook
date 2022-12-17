//	Compare rough execution times of examples from Chapter 1.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	//	Output rough execution time for string concatenation.
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed.Milliseconds())

	//	Output rough execution time for strings.Join().
	start = time.Now()
	s = strings.Join(os.Args[1:], " ")
	elapsed = time.Since(start)
	fmt.Println(elapsed.Milliseconds())
}

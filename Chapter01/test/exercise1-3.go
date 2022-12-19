//	Compare execution times of examples from Chapter 1.
//	TODO: get test args/files to benchmark function.
package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

func StringJoinSlow() {
	os.Args = append(os.Args, "here are some args")
	start := time.Now()
	s, sep := "", ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	elapsed := time.Since(start)
	fmt.Println(elapsed.Milliseconds())
}

func StringJoinFast() {
	os.Args = append(os.Args, "args.txt")
	start := time.Now()
	s := strings.Join(os.Args[1:], " ")
	fmt.Println(s)
	elapsed := time.Since(start)

	fmt.Println(elapsed.Milliseconds())
}

func BenchmarkStringSplit(b *testing.B) {
	StringJoinSlow()
	StringJoinFast()
}

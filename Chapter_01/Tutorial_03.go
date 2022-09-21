// 1.3. Finding Duplicate Lines

// Programs for file copying, printing, searching, sorting, counting, and the like all have a similar
// structure: a loop over the input, some computation on each element, and generation of output
// on the fly or at the end. We’ll show three variants of a program called dup; it is partly inspired
// by the Unix uniq command, which looks for adjacent duplicate lines. The structures and
// packages used are models that can be easily adapted.
// The first version of dup prints each line that appears more than once in the standard input,
// preceded by its count. This program introduces the if statement, the map data type, and the
// bufio package.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

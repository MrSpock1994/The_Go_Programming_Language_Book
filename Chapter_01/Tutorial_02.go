// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

// 1.2. Command-Line Arguments

// Most programs process some input to produce some output; that’s pretty much the definition
// of computing. But how does a program get input data on which to operate? Some programs
// generate their own data, but more often, input comes from an external source: a file, a network
// connection, the output of another program, a user at a keyboard, command-line arguments,
// or the like. The next few examples will discuss some of these alternatives, starting with com-
// mand-line arguments.
// The os package provides functions and other values for dealing with the operating system in a
// platform-independent fashion. Command-line arguments are available to a program in a
// variable named Args that is part of the os package; thus its name anywhere outside the os
// package is os.Args.
// The variable os.Args is a slice of strings. Slices are a fundamental notion in Go, and we’ll talk
// a lot more about them soon. For now, think of a slice as a dynamically sized sequence s of
// array elements where individual elements can be accessed as s[i] and a contiguous subse-
// quence as s[m:n]. The number of elements is given by len(s). As in most other program-
// ming languages, all indexing in Go uses half-open intervals that include the first index but
// exclude the last, because it simplifies logic. For example, the slice s[m:n], where 0 ≤ m ≤ n ≤
// len(s), contains n-m elements.
// The first element of os.Args, os.Args[0], is the name of the command itself; the other ele-
// ments are the arguments that were presented to the program when it started execution. A
// slice expression of the form s[m:n] yields a slice that refers to elements m through n-1, so the
// elements we need for our next example are those in the slice os.Args[1:len(os.Args)]. If m
// or n is omitted, it defaults to 0 or len(s) respectively, so we can abbreviate the desired slice as
// os.Args[1:].

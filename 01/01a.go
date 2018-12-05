// Advent of Code 2018 - Day 1, Part 1
// n0j

package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {

	// usage
	if len(os.Args) != 2 {
		nope(errors.New("<infile>"))
	}

	// read input file 
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		nope(err)
	}
	ops := strings.Split(string(data), "\n")

	// input file may contain trailing newline, remove
	if ops[len(ops)-1] == "" {
    	ops = ops[:len(ops)-1]
	}

	// the meat
	// x: sum from previous iteration
	// y: sum this iteration
	// n: num value from input
	var x, y int
	for i, op := range ops {
		
		// get input value
		n, err := strconv.Atoi(string(op[1:]))
		if err != nil {
			nope(err)
		}
		
		// ascii(43) = '+'
		if op[0] == 43 {
			y = x + n
			fmt.Printf("[%d] %d + %d = %d\n", i, x, n, y)

		// ascii(45) = '-'
		} else if op[0] == 45 {
			y = x - n
			fmt.Printf("[%d] %d - %d = %d\n", i, x, n, y)

		// bad op
		} else {
			nope(errors.New("input operation must be + or -"))
		}

		// prep next iteration
		x = y
	}

	// result
	fmt.Printf("[!] TOTAL: %d\n", x)
}

func nope(err error) {
	fmt.Fprintf(os.Stderr, "%v: %v\n", os.Args[0], err)
	os.Exit(1)
}
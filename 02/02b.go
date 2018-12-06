// Advent of Code 2018 - Day 2, Part 2
// n0j
//
// "Where cleverness is lackin, brute force is crackin." -Me, 2 mins ago

package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
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
	ids := strings.Split(string(data), "\n")

	// input file may contain trailing newline, remove
	if ids[len(ids)-1] == "" {
    	ids = ids[:len(ids)-1]
	}

	// the meat
	low := len(ids[0])+1 	// current lowest edit distance
	var edit input			// per-iteration edit distance
	var sol strings			// current solution string, repeat chars removed
	for i := 0; i < len(ids)-1; i++ { // for each string
		for j := i+1; j < len(ids); j++ { // for each other string
			
			// scan string pair for number of disimilar chars
			edit = 0
			for k := 0; k < len(ids[0]); k++ {
				if ids[i][k] != ids[j][k] {
					edit += 1
				}
			}

			// compare to previous low edit, store as sol if lower
			if edit < low {
				low = edit
				sol = ""
				for k := 0; k < len(ids[0]); k++ {
					if ids[i][k] == ids[j][k] {
						sol = sol + string(ids[i][k])
					}
				}
			}
		}
	}

	// print solution, return sledgehammer to shed in shame, take up quilting
	fmt.Println(sol)
}

func nope(err error) {
	fmt.Fprintf(os.Stderr, "%v: %v\n", os.Args[0], err)
	os.Exit(1)
}
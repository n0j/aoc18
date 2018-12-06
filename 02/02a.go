// Advent of Code 2018 - Day 2, Part 1
// n0j
//
// worked for AoC input but would not in general case as quads+
// would count for trips (thanks AV-IO) - oops

package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
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
	var dubs, trips int
	for _, id := range ids {

		// string -> slice of runes
		r := []rune(id)

		// sort tip from https://siongui.github.io/2017/05/07/go-sort-string-slice-of-rune/
		sort.Slice(r, func(i, j int) bool {
			return r[i] < r[j]
		})
		r = append(r, 0x00) // remove last-pair exception with trailing null
		
		// scan for dubs, trips
		i := 0
		var dub, trip bool
		for i < len(r)-2 {
			if r[i] == r[i+1] {
				if r[i] == r[i+2] {
					trip = true
					i += 1	// skip ahead to avoid counting dub at end of trip
				} else {
					dub = true
				}
			}
			i += 1
		}

		// if found this iteration, increment total count
		if dub { 
			dubs += 1 
		}
		if trip {
			trips += 1
		}

	}

	fmt.Printf("[!] %d dubs %d trips: checksum %d\n", dubs, trips, dubs*trips)
}

func nope(err error) {
	fmt.Fprintf(os.Stderr, "%v: %v\n", os.Args[0], err)
	os.Exit(1)
}
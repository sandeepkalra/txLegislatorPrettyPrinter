package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
	"strings"
)

var (
	in          = os.Getenv("IN")
	sectionName = os.Getenv("SECT")
)

// algo:
// search for patterns. if they are not found, we simply print back
// if patterns are available, put all name, and the location.
//
//	-- sort by location to get unique patterns
//	-- note all unique pattern means the previous line has (x) entry. x can be 0-9,A-Z,a-z
//	-- then look for next such unique pattern
//	-- repeat and see previous line for (x).
//
// once we have all unique-patterns, and the order, we then tab to print pretty
func prettySection(in string) string {
	found := false
	const max = math.MaxInt32
	exprs := map[string]int{
		`:\n\([a-z]-[0-9]\) `: max,
		`:\n\([A-Z]\) `:       max,
		`:\n\([A-Z]-[0-9]\) `: max,
		`:\n\([a-z])\) `:      max,
		`:\n\([0-9]\) `:       max,
		`:\n\([0-9]-[a-z]\) `: max,
		`:\n\([0-9]-[A-Z]\) `: max,
		`:\n\([i|x|v]+\) `:    max,
	}
	for exp := range exprs {
		b, e := regexp.MatchString(exp, in)
		if e == nil {
			if b {
				smallestLoc := max
				r, e := regexp.Compile(exp)
				if e != nil {
					fmt.Println("ERROR ", e)
					continue
				}
				locations := r.FindIndex([]byte(in))
				for _, loc := range locations {
					if loc < smallestLoc {
						smallestLoc = loc
					}
				}
				exprs[exp] = smallestLoc
				found = true
			}
		}
	}
	if !found {
		fmt.Println(in)
	} else {
		newExpr := make(map[string]int)
		locations := []int{}
		for _, l := range exprs {
			if l != max {
				locations = append(locations, l)
			}
		}
		sort.Slice(locations, func(i, j int) bool {
			return locations[i] < locations[j]
		})
		/*
			for _, l := range locations {
				fmt.Println("debug:", l)
			}
			for xpr, vLoc := range exprs {
				fmt.Println("debug:", xpr, vLoc)
			}
		*/
		for k, lo := range locations {
			for exp, vLoc := range exprs {
				if vLoc == lo && exp != "" {
					//				fmt.Println("debug: ", exp, lo)
					if arr := strings.Split(exp, `:\n`); len(arr) > 1 {
						ss := arr[1]
						newExpr[ss] = k + 1 // k is number of spaces that are needed
					} else {
						fmt.Println("error: ", arr)
					}
				}
			}
		}
		/*
			for xpr, v := range newExpr {
				fmt.Println("found:", xpr, "space :", v)
			}
		*/
		inp := strings.Split(in, "\n")

		for _, l := range inp {
			printed := false
			for xpr, spc := range newExpr {
				s := ""
				for i := 0; i < spc; i++ {
					s += "  "
				}
				r, _ := regexp.Compile(xpr)
				if r.Match([]byte(l)) {
					fmt.Printf("%v%v\n", s, l)
					printed = true
				}
			}
			if !printed {
				fmt.Println(l)
			}
		}
	}
	return in
}

func main() {
	if in == "" {
		in = "tx1101"
	}
	if sectionName == "" {
		sectionName = "Sec. 1101."
	}
	inf, e := os.ReadFile(in + ".txt")
	if e != nil {
		log.Fatal("failed to read file")
	}
	inps := string(inf)
	chapters := strings.Split(inps, "SUBCHAPTER")
	for _, chapContent := range chapters {
		sections := strings.Split(chapContent, sectionName)
		for _, ssectionName)ecCont := range sections {
			sectionFull := sectionName
			fmt.Print("\n")
			sectLines := strings.Split(secCont, "\n")
			for _, line := range sectLines {
				l := strings.TrimSpace(line)
				if _, b := strings.CutPrefix(l, "Added by Acts 2"); b {
					continue
				}
				if _, b := strings.CutPrefix(l, "Acts 2"); b {
					continue
				}
				if _, b := strings.CutPrefix(l, "Amended by:"); b {
					continue
				}
				if l == "" {
					continue
				}
				m := l + "\n"
				sectionFull += m
			}
			prettySection(sectionFull)
		} // section complete

	}
}

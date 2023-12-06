package main

import (
	"AoC/23/utils"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var dict = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

// Checks if our stored indexes contains the given set of indexes
func contains(s [][]int, idx []int) bool {
	for _, v := range s {
		if v[0] == idx[0] && v[1] == idx[1] {
			return true
		}
	}
	return false
}

func parse(s []string) []string {
	ret := []string{}

	for _, l := range s {
		cp := l
		matches := [][]int{}

		// Iterate over every character of the string
		// Overkill but this makes sure we dont skip any possible match
		// We are thus working with an incomplete string, messing up the indexes. That is why we later add +i when we create our match slice
		for i := range l {
			// Iterate over each key of our dict to check if there is a match
			for k := range dict {
				idx := strings.Index(l[i:], k)

				if idx != -1 {
					match := []int{idx + i, idx + len(k) + i}

					if !contains(matches, match) {
						matches = append(matches, match)
					}
				}
			}
		}

		for j, idx := range matches {
			start, end := idx[0], idx[1]

			res := l[start:end]
			num := dict[res]
			if num != 0 {
				// +i to account for the added number.
				// Wont run properly if num != 0 fails and no new number is added to our cp string, but should not happen here
				cp = cp[:start+j] + fmt.Sprint(num) + cp[start+j:]
			}
		}
		ret = append(ret, cp)
	}

	return ret
}

func calibrate(s []string) int {
	ret := 0

	for _, str := range s {
		first := 0
		second := 0
		for _, c := range str {
			if unicode.IsNumber(c) {
				i, _ := strconv.Atoi(string(c))
				if first == 0 {
					first = i
				} else {
					second = i
				}
			}
		}
		// If only one number in string
		if second == 0 {
			second = first
		}
		ret += first*10 + second
	}

	return ret
}

var input = utils.Input("input.txt")

// var test = utils.Input("test.txt")

func part1() {
	fmt.Println("Part 1")
	fmt.Println(calibrate(input))
}

func part2() {
	fmt.Println("----------------")
	fmt.Println("Part 2")
	fmt.Println(calibrate(parse(input)))

}

func main() {
	part1()
	part2()
}

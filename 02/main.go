package main

import (
	"AoC/23/utils"
	"fmt"
	"strconv"
	"strings"
)

const Red, Green, Blue = 12, 13, 14

var input = utils.Input("input.txt")

// var test = utils.Input("test.txt")

func parse(s []string) [][]map[string]int {
	ret := [][]map[string]int{}

	for _, l := range s {
		line := []map[string]int{}

		cp := l
		idx := strings.Index(cp, ": ")
		cp = l[idx+2:]
		sets := strings.Split(cp, "; ")

		for _, v := range sets {
			tmp := strings.Split(v, ", ")
			m := make(map[string]int)

			for _, k := range tmp {
				split := strings.Split(k, " ")
				num, _ := strconv.Atoi(split[0])
				m[split[1]] = num
			}

			line = append(line, m)
		}

		ret = append(ret, line)
	}

	return ret
}

func calc(s [][]map[string]int) int {
	var ret, sum int

	for i, game := range s {
		var red, green, blue int
		possible := true

		for _, set := range game {
			for k, v := range set {
				switch k {
				case "red":
					if v > Red {
						possible = false
					}
					if v > red {
						red = v
					}
				case "green":
					if v > Green {
						possible = false
					}
					if v > green {
						green = v
					}
				case "blue":
					if v > Blue {
						possible = false
					}
					if v > blue {
						blue = v
					}
				}
			}
		}
		if possible {
			ret += i + 1
		}
		sum += red * green * blue
	}

	fmt.Println("Part 1 & 2")
	fmt.Println("Number of possible games :", ret)
	fmt.Println("Sum of minimum poswer for all games:", sum)

	return ret
}

func main() {
	calc(parse(input))
}

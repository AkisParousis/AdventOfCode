package main

import (
	"fmt"
	_ "embed"
	"strconv"
	"strings"
)

const initial_dial int = 50

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func dial(direction string, degrees int, dial int) (int, int) {
    var quotient int

    switch direction {
    case "R":
        quotient = (dial + degrees) / 100
        dial = (dial + degrees) % 100
    case "L":
		if dial - degrees <= 0 {
			if (degrees - dial) / 100 == 0 {
				if dial != 0 {
					quotient = 1
				}
			} else {
				quotient = (degrees - dial) / 100 + 1
				if dial == 0 {
					quotient -= 1
				}
			}
		}
        dial = (dial - degrees + (quotient+1) * 100) % 100
    default:
        return -1, -1
    }
	// fmt.Printf("Dial turned %s by %d degrees: new dial position %d, crossed zero %d times\n", direction, degrees, dial, quotient)
    return dial, quotient
}
	
//go:embed files/input.txt
var inputFile string

func main() {
    dialValue := initial_dial
	quotient := 0
	count := 0
	total_count := 0
	counts := make(map[int]int)

	for _, line := range strings.Split(strings.TrimSpace(inputFile), "\n") {
		// fmt.Println(line)
		direction := string(line[0])
		degrees, _ := strconv.Atoi(line[1:])
		dialValue, quotient = dial(direction, degrees, dialValue)
		counts[quotient]++
		if dialValue == 0 {
			count += 1
		}
		total_count += quotient
	}

	recomputed := 0
	for k, v := range counts {
		recomputed += k * v
	}
	fmt.Printf("Times dial landed on zero: %d\n", count)
	fmt.Println("recomputed (from histogram):", recomputed)
}
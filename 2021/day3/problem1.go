package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	c := [12]int{}
	gamma := 0.0
	epsilon := 0.0

	for scanner.Scan() {
		line := scanner.Text()
		for i, val := range line {
			if val == 48 {
				c[i] += -1
			} else {
				c[i] += 1
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(c)
	fmt.Println(len(c))
	g := [12]int{}
	e := [12]int{}

	for i, num := range c {
		if num > 0 {
			g[i] = 1
			e[i] = 0
			gamma += math.Pow(2, float64(len(c)-i-1))
		} else {
			g[i] = 0
			e[i] = 1
			epsilon += math.Pow(2, float64(len(c)-i-1))
		}
	}
	fmt.Println(g)
	fmt.Println(e)
	fmt.Printf("Gamma Rate: %f\nEpsilon Rate: %f\nPower: %f\n", gamma, epsilon, gamma*epsilon)
}

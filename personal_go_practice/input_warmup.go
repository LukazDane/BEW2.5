package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	yearCalc()
}

func yearCalc() int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your age: ")
	age, _ := reader.ReadString('\n')
	age = strings.TrimSuffix(age, "\n")
	// fmt.Print("Your age ", age)

	year := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your birth year: ")
	by, _ := year.ReadString('\n')
	by = strings.TrimSuffix(by, "\n")
	// fmt.Print("Your year ", by)

	yearint, _ := strconv.Atoi(by)
	ageint, _ := strconv.Atoi(age)
	current := yearint + ageint
	fmt.Printf("%d", current)
	return current
}

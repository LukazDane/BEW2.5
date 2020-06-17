package main

import (
	"bufio"
	"fmt"
	"os"
)

type House struct {
	numberOfRooms string
	city          string
	address       string
	price         string
}

func main() {
	answer := true
	if answer {
		house()
	}
}

func house() {
	var ahouse House
	rooms := bufio.NewReader(os.Stdin)
	fmt.Print("How many rooms? ")
	ahouse.numberOfRooms, _ = rooms.ReadString('\n')
	// fmt.Print("Number of Rooms: ", ahouse.numberOfRooms)
	loc := bufio.NewReader(os.Stdin)
	fmt.Print("City? ")
	ahouse.city, _ = loc.ReadString('\n')

	add := bufio.NewReader(os.Stdin)
	fmt.Print("Address? ")
	ahouse.address, _ = add.ReadString('\n')

	cost := bufio.NewReader(os.Stdin)
	fmt.Print("Price? ")
	ahouse.price, _ = cost.ReadString('\n')

	fmt.Print(ahouse)

	// ans := bufio.NewReader(os.Stdin)
	// fmt.Print("add another? ")
	// if ans == "n"  {
	// 	answer = ans.ReadString('\n')
	// }
}

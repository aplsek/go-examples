package main


import (
	"fmt"
	//"math/rand"
)

func move(from int, to int) {
	fmt.Printf("%d -> %d \n", from, to)
}

func getMiddlePeg(peg1 int, peg2 int) int {
	switch peg1 {
		case 1:
			switch peg2 {
			case 2 :
				return 3
			case 3 :
				return 2
			}
		case 2: 
			switch peg2 {
			case 1:
				return 3
			case 3:
				return 1 
			}
		case 3:
			switch peg2 {
			case 2:
				return 1
			case 1:
				return 2 
			}
		default: fmt.Println("default")
	}
	return 1
}

func hanoi(nDisks int, peg1 int, peg2 int)  {
	if nDisks == 1 {
		move (peg1, peg2)
		return
	}
	var peg3 = getMiddlePeg(peg1, peg2)

	hanoi(nDisks -1 , peg1, peg3)

	move(peg1, peg2)

	hanoi(nDisks -1 , peg3, peg2)
	
	return
}



func main() {
	hanoi(3,1,3)
	//fmt.Printf(" Hanoi solution = %s\n", res)
}

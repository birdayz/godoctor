package main

import "fmt"

func main() {
	x := []int{5, 3, 2, 1, 8}
	fmt.Println("please enter a selection:")
	//x = input from user
	for _, sf := range x {
		dv := sf
		switch sf {
		case 1:

		case 2:
			i := new(int)
			if dv != 0 { // <<<<< var,16,7,16,9,newVar,pass
				*i = dv
			}
		}

	}
}

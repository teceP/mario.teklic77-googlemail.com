package utils

import "fmt"

func ReadInput() string {
	var input string
	fmt.Scanln(&input)
	return input
}

func CheckInputLength(input string) bool {

	if len(input) != 3 {
		fmt.Println("Wrong input. Example: A/1")
		return false
	}

	//if input[0] == letter, if input[2] == int -> false
	return true
}

func ConvertInput(input string) (int, int) {
	var col, row = -1, -1

	for _, c := range input {
		switch c {

		case 'A':
			col = 0
		case 'B':
			col = 1
		case 'C':
			col = 2
		case '/':
			//nothing
		case '1':
			row = 0
		case '2':
			row = 1
		case '3':
			row = 2
		}
	}

	return col, row
}

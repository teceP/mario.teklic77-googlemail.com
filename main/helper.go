package main

import "fmt"

//Reads the next line of the users cmd-line-input
func ReadInput() string {
	var input string
	fmt.Scanln(&input)
	return input
}

//Checks if the string length is 2
func CheckInputLength(input string) bool {

	if len(input) != 2 {
		fmt.Println("Wrong input. Example: A1")
		return false
	}

	//if input[0] == letter, if input[2] == int -> false
	return true
}

//Converts the user-input into 2 coordinates (x/y)
//Returns -1 (x and/or y) if input was false
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

//Evaluates the users choice of playing against a computer or a human
//Returns an error if the input wasnt correct
func EvaluateUserDec(userDec string) (bool, error) {
	if userDec == "Y" {
		fmt.Println("You will play against the computer!")
		return true, nil
	} else if userDec == "n" {
		fmt.Println("You will play against another human.")
		return false, nil
	} else {
		return false, fmt.Errorf("Cannot handle user input.")
	}
}

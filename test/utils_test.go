package test

import (
	"fmt"
	"testing"
	"tictactoe_go/utils"
)

func Test_CheckInputLength_good(t *testing.T) {
	legalInput := utils.CheckInputLength("A/1")

	fmt.Println(legalInput)

	if legalInput != true {
		t.Errorf("Result was: %t, want: %t.", legalInput, true)
	}
}

func Test_CheckInputLength_bad(t *testing.T) {
	legalInput := utils.CheckInputLength("A1")

	fmt.Println(legalInput)

	if legalInput != false {
		t.Errorf("Result was: %t, want: %t.", legalInput, false)
	}
}

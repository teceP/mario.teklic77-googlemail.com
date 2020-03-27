package test

//UNIT TESTS

import (
	"testing"
	"tictactoe_go/utils"
)

func Test_CheckInputLength_good(t *testing.T) {
	legalInput := utils.CheckInputLength("A/1")

	if legalInput != true {
		t.Errorf("Result was: %t, want: %t.", legalInput, true)
	}
}

func Test_CheckInputLength_bad(t *testing.T) {
	legalInput := utils.CheckInputLength("A1")

	if legalInput != false {
		t.Errorf("Result was: %t, want: %t.", legalInput, false)
	}
}

func Test_ConvertInput_good_0(t *testing.T) {
	x, y := utils.ConvertInput("A/1")

	if x != 0 {
		t.Errorf("Result was: %v, want: %v.", x, 0)
	}

	if y != 0 {
		t.Errorf("Result was: %v, want: %v.", y, 0)
	}
}

func Test_ConvertInput_good_1(t *testing.T) {
	x, y := utils.ConvertInput("B/2")

	if x != 1 {
		t.Errorf("Result was: %v, want: %v.", x, 1)
	}

	if y != 1 {
		t.Errorf("Result was: %v, want: %v.", y, 1)
	}
}

func Test_ConvertInput_good_2(t *testing.T) {
	x, y := utils.ConvertInput("C/3")

	if x != 2 {
		t.Errorf("Result was: %v, want: %v.", x, 2)
	}

	if y != 2 {
		t.Errorf("Result was: %v, want: %v.", y, 2)
	}
}

func Test_ConvertInput_bad_0(t *testing.T) {
	x, y := utils.ConvertInput("D/4")

	if x != -1 {
		t.Errorf("Result was: %v, want: %v.", x, -1)
	}

	if y != -1 {
		t.Errorf("Result was: %v, want: %v.", y, -1)
	}
}

func Test_ConvertInput_bad_1(t *testing.T) {
	x, y := utils.ConvertInput("C/5")

	if x != 2 {
		t.Errorf("Result was: %v, want: %v.", x, 2)
	}

	if y != -1 {
		t.Errorf("Result was: %v, want: %v.", y, -1)
	}
}

func Test_ConvertInput_bad_2(t *testing.T) {
	x, y := utils.ConvertInput("F/3")

	if x != -1 {
		t.Errorf("Result was: %v, want: %v.", x, -1)
	}

	if y != 2 {
		t.Errorf("Result was: %v, want: %v.", y, 2)
	}
}

func Test_EvaluateUserDec_good_0(t *testing.T) {
	b, err := utils.EvaluateUserDec("Y")

	if b != true {
		t.Errorf("Result was: %t, want: %t.", b, true)
	}

	if err != nil {
		t.Error("Result was: ", err, " want: ", nil)
	}
}

func Test_EvaluateUserDec_good_1(t *testing.T) {
	b, err := utils.EvaluateUserDec("n")

	if b != false {
		t.Errorf("Result was: %t, want: %t.", b, false)
	}

	if err != nil {
		t.Error("Result was: ", err, " want: ", nil)
	}
}

func Test_EvaluateUserDec_bad_0(t *testing.T) {
	b, err := utils.EvaluateUserDec("/")

	if b != false {
		t.Errorf("Result was: %t, want: %t.", b, false)
	}

	if err == nil {
		t.Error("Result was: ", err, " want: ", nil)
	}
}

func Test_EvaluateUserDec_bad_1(t *testing.T) {
	b, err := utils.EvaluateUserDec("q")

	if b != false {
		t.Errorf("Result was: %t, want: %t.", b, false)
	}

	if err == nil {
		t.Error("Result was: ", err, " want: ", nil)
	}
}

func Test_EvaluateUserDec_bad_2(t *testing.T) {
	b, err := utils.EvaluateUserDec("1")

	if b != false {
		t.Errorf("Result was: %t, want: %t.", b, false)
	}

	if err == nil {
		t.Error("Result was: ", err, " want: ", nil)
	}
}

func Test_EvaluateUserDec_bad_3(t *testing.T) {
	b, err := utils.EvaluateUserDec("Yn")

	if b != false {
		t.Errorf("Result was: %t, want: %t.", b, false)
	}

	if err == nil {
		t.Error("Result was: ", err, " want: ", nil)
	}
}

func Test_EvaluateUserDec_bad_4(t *testing.T) {
	b, err := utils.EvaluateUserDec("nY")

	if b != false {
		t.Errorf("Result was: %t, want: %t.", b, false)
	}

	if err == nil {
		t.Error("Result was: ", err, " want: ", nil)
	}
}

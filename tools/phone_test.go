package tools

import "testing"

func TestPhoneNumber(t *testing.T){
	t.Log(ValidatePhone("86000"))
	t.Log(ValidatePhone("10000"))
	t.Log(ValidatePhone("8613041259681"))
}

func TestPhoneNumber2(t *testing.T){
	t.Log(ValidatePhone("+86000"))
	t.Log(ValidatePhone("+10000"))
	t.Log(ValidatePhone("+8613041259681"))
	t.Log(ValidatePhone("+86-13041259681"))
	t.Log(ValidatePhone("+86 -13041259681  "))
	t.Log(ValidatePhone("+8623041259681"))
}
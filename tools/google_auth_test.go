package tools

import "testing"

func TestGetGoogleCode(t *testing.T){
	screctCode := "PNCSSCERENJIYBDW"
	offset := 25
	t.Log(getCode(screctCode, int64(offset)))
}


func TestGetGoogleScrectCode(t *testing.T){
	t.Log(GetSecret())
}


func TestValidateGoogleCode(t *testing.T){
	screctCode := "PNCSSCERENJIYBDW"
	offset := 5
	code := getCode(screctCode, int64(offset))
    t.Log("code: ", code)
	t.Log(VerifyCode(screctCode, code))
}

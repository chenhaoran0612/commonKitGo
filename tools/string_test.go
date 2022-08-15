package tools

import "testing"

func TestFirstToUpper(t *testing.T){
	t.Log(FirstToUpper("abcee"))
	t.Log(FirstToUpper("1bcee"))
	t.Log(FirstToUpper("Abcee"))
	t.Log(FirstToUpper(""))
	t.Log(FirstToUpper("-1"))
}

func TestFirstToLower(t *testing.T){
	t.Log(FirstToLower("abcee"))
	t.Log(FirstToLower("1bcee"))
	t.Log(FirstToLower("Abcee"))
	t.Log(FirstToLower(""))
	t.Log(FirstToLower("-1"))
}

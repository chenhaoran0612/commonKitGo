package tools

import (
	"testing"
)


func TestIsNum(t *testing.T){
	t.Log(IsNum("1222"))
	t.Log(IsNum("0"))
	t.Log(IsNum("0a"))
	t.Log(IsNum("a"))
	t.Log(IsNum("0000"))
	t.Log(IsNum("a000"))
	t.Log(IsNum(".0"))
	t.Log(IsNum("-"))
	t.Log(IsNum("0.9"))
}

func TestIsMinerName(t *testing.T){
	t.Log(IsMinerName("1222"))
	t.Log(IsMinerName("0.ddde._3223a"))
	t.Log(IsMinerName("0a"))
	t.Log(IsMinerName("abcd"))
}

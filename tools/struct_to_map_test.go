package tools

import "testing"

func TestStructToMap(t *testing.T){
	type A struct{
		A int
	}

	t.Log(StructToMap(&A{A: 1}, "json"))
	t.Log(StructToMap(&A{A: 1}, ""))
}


func TestStructToMap2(t *testing.T){
	type A struct{
		A int `json:"a"`
	}

	t.Log(StructToMap(&A{A: 1}, "json"))
	t.Log(StructToMap(&A{A: 1}, ""))
}

func TestStructToMap3(t *testing.T){
	type A struct{
		A int `json:"a"`
		B int `json:"b"`
	}

	t.Log(StructToMap(&A{A: 1}, "json"))
	t.Log(StructToMap(A{A: 1}, ""))
}

func TestStructToMap4(t *testing.T){

	t.Log(StructToMap(struct {}{}, "json"))
}

func TestStructToMap5(t *testing.T){

	// 抛错误
	t.Log(StructToMap(nil, "json"))
}

func TestStructToMap6(t *testing.T){

	// 抛错误
	t.Log(StructToMap([]int{}, "json"))
}

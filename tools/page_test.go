package tools

import "testing"

func TestPage(t *testing.T){
	start, end := PageOffset(1, 20, 50)
	t.Log(start, end)

	start, end = PageOffset(2, 20, 50)
	t.Log(start, end)

	start, end = PageOffset(3, 20, 50)
	t.Log(start, end)

	start, end = PageOffset(-1, -20, 50)
	t.Log(start, end)
}

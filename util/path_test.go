package util

import (
	"testing"
	"strconv"
)

func TestAppend(t *testing.T) {
	//root := "http://localhost:8081\\"
	//root2 := "http://localhost:8081"
	//tail := "asdf/asdf\\as"
	//tail2 := "/asldf/asdf/fdas"
	//t.Log(PathAppend(root, tail))
	//t.Log(PathAppend(root, tail2))
	//t.Log(PathAppend(root2, tail))
	//t.Log(PathAppend(root2, tail2))
	//t.Log(test())
	num,err := strconv.Atoi("")
	if err!=nil{
		t.Error(err)
	}else{
		t.Log(num)
	}
}

func test() (path string){
	root := "http://localhost:8081\\"
	tail := "a/b\\c"
	tail2 := "/d/e/f"
	tail3 := "1/2/3"
	tail4 := "/4/5/6/"
	tail5 := "7/8/9/"
	tail6 := ""
	path = PathAppend(root, tail, tail2, tail3, tail6, tail4, tail5, tail6)
	return path
}

func TestPathAppend2(t *testing.T) {
	root := "http://localhost:8081\\"
	tail := "a/b\\c"
	tail2 := "/d/e/f"
	tail3 := "1/2/3"
	tail4 := "/4/5/6/"
	tail5 := "7/8/9/"
	tail6 := ""
	t.Log(PathAppend(root, tail, tail2, tail3, tail6, tail4, tail5, tail6))
}


func BenchmarkPathAppend2(b *testing.B) {
	root := "http://localhost:8081\\"
	tail := "a/b\\c"
	tail2 := "/d/e/f"
	tail3 := "1/2/3"
	tail4 := "/4/5/6/"
	tail5 := "7/8/9/"
	tail6 := ""
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		//PathAppend(root, tail)
		PathAppend(root, tail, tail2, tail3, tail6, tail4, tail5, tail6)
	}
}
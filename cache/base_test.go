package cache

import "testing"

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func TestSet(t *testing.T) {
	//err := Set("t1", "mimi", 0)
	//checkErr(err)
	ti, err := GetExpire("t1")
	checkErr(err)
	t.Log(ti)
	v, err := Get("t1")
	checkErr(err)
	t.Log(v)
}

package util

import (
	"testing"
)

func TestIsArrEmpty(t *testing.T) {
	t1 := []string{}
	var t2 []string
	t3 := []string{"sdf"}
	if !IsStringArrEmpty(t1) {
		t.Error("t1")
	}
	if !IsStringArrEmpty(t2) {
		t.Error("t2")
	}
	if IsStringArrEmpty(t3) {
		t.Error("t3")
	}
}

func TestStringArrDelete(t *testing.T) {
	arr1 := []string{"a", "b", "c"}
	arr2 := StringArrDelete(arr1, "b")
	if len(arr2) != 2 || arr2[0] != "a" || arr2[1] != "c" {
		t.Error(arr1, arr2)
	}
}

func TestArrSlice(t *testing.T) {
	arr1 := []string{"a", "b", "c"}
	slice1 := make([]string, len(arr1), len(arr1))
	copy(slice1, arr1)
	t.Log("slice1:", slice1)
	slice1[0] = "d"
	t.Log("slice1:", slice1)
	t.Log("arr1:", arr1)
}

func TestArrInterface(t *testing.T) {
	list := testReturnIntarfaceArr()
	aa := list.([]*TempObj)
	t.Log(aa[0].Id)
}

func testReturnIntarfaceArr() interface{} {
	list := make([]*TempObj, 0, 10)
	list = append(list, &TempObj{"a"})
	return list
}

type TempObj struct {
	Id string
}
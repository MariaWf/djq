package util

import (
	"testing"
	"time"
)

func TestParseTimeFromDB(t *testing.T) {
	str := StringTime4DB(time.Now().Add(time.Hour * 10))
	v, err := ParseTimeFromDB(str)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}

	k := true
	testBool(&k)
	t.Log(k)
}

func testBool(test *bool) {
	*test = false
}

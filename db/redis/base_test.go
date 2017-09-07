package redis

import (
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	conn := Get()
	str, err := conn.Get("name1").Result()
	if err != nil {
		t.Error(err)
	} else if str != "mimiwawa"{
		t.Error(str)
	}else{
		t.Log(str)
	}
}

func TestSet(t *testing.T) {
	conn := Get()
	str, err := conn.Set("name1", "mimiwawa", time.Minute * 1).Result()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(str)
	}
}
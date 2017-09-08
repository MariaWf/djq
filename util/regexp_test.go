package util

import "testing"

func TestMatchDescription(t *testing.T) {
	maps := make(map[string]bool)
	maps["a12aSdf_收到风？》《“：{Z}XCPFD{S{P4\nsdf"] = true
	maps["012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567892"] = false
	for k, v := range maps {
		if MatchDescription(k) != v {
			t.Errorf("%v应为%v", k, v)
		}
	}
}

func TestMatchName(t *testing.T) {
	maps := make(map[string]bool)
	maps["a12aSdf_sdf"] = true
	maps["ass"] = false
	maps["name0"] = true
	maps["12asldfkj"] = false
	maps["fd1dsafdesfd1dsafdesfd1dsafdes21"] = true
	maps["fd1dsafdesfd1dsafdesfd1dsafdes21d"] = false
	for k, v := range maps {
		if MatchName(k) != v {
			t.Errorf("%v应为%v", k, v)
			goto test
		}
	}
	t.Log("test1")
	return

	test:
	t.Log("test")
}

func TestMatchMobile(t *testing.T) {
	maps := make(map[string]bool)
	maps["12345678901"] = true
	maps["ass"] = false
	maps["12asldfkj"] = false
	maps["123456789012"] = false
	maps["22345678901"] = false
	for k, v := range maps {
		if MatchMobile(k) != v {
			t.Errorf("%v应为%v", k, v)
		}
	}
}
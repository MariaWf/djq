package constant

import (
	"fmt"
	"mimi/djq/config"
	"testing"
)

func BenchmarkConfig(b *testing.B) {
	b.ResetTimer()
	b.StartTimer()
	//Config = config.GetInstance()
	for i := 0; i < b.N; i++ {
		//if Config.Mymap["adminName"] != "admin" {
		//	fmt.Println("err")
		//}
		if config.Get("adminName") != "admin" {
			fmt.Println("err")
		}
	}
}

func TestApiType(t *testing.T) {
	t.Log(ApiTypeMi, ApiTypeUi, ApiTypeSi, ApiTypeOpen)
}

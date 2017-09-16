package config

import (
	"testing"
)

func TestGet(t *testing.T) {
	t.Log(Get("wxpay_sandboxnew_131"))
}

func BenchmarkGet(b *testing.B) {
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Get("mysqlDataSourceName")
	}
}

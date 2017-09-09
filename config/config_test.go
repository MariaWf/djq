package config

import (
	"testing"
)

func TestGet(t *testing.T) {
	t.Log(Get("mysqlDataSourceName"))
}

func BenchmarkGet(b *testing.B) {
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Get("mysqlDataSourceName")
	}
}

package handler

import (
	"testing"
	"mimi/djq/service"
)

func BenchmarkGetRoleServcieInstance(b *testing.B) {
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		&service.Role{}
	}
}

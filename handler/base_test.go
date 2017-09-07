package handler

import "testing"

func BenchmarkGetRoleServcieInstance(b *testing.B) {
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		GetRoleServiceInstance()
	}
}

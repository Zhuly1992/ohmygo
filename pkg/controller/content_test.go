package controller

import "testing"

//func Test_ForDemo(t *testing.T) {
//	ForDemo()
//}

func BenchmarkForDemo(b *testing.B) {
	b.StopTimer()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ForDemo()

	}
}

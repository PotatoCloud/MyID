package MyID

import "testing"

func TestNew(t *testing.T) {
	myid := New(100)
	t.Logf("id: %d", myid.Generate())
}

func BenchmarkNew(b *testing.B) {
	myid := New(100)
	for i := 0; i < b.N; i++ {
		myid.Generate()
	}
}

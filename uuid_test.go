package uuid

import (
	"testing"
	"strconv"
)

func dedup(x [][]byte) [][]byte {
	visited := map[string]bool{}
	r := [][]byte{}

	for v := range x {
		w := string(x[v])
		if !visited[w] {
			visited[w] = true
			r = append(r, x[v])
		}
	}
	return r
}

func TestGenerate(t *testing.T) {
	s := make([][]byte, 1000)

	for n := 0; n < 1000; n++ {
		u, err := Generate()
		if err != nil {
			t.Fatalf(err.Error())
		}
		s[n] = u
	}

	if len(dedup(s)) != len(s) {
		t.Errorf("duplicates generated : Set-%d Seq-%d", len(dedup(s)), len(s))
	}
}

func TestVersion(t *testing.T) {
	u, err := Generate()
	if err != nil {
		t.Fatalf(err.Error())
	}
	version, err := strconv.ParseUint(string(u[14]) + string(u[15]), 16, 64)
	if err != nil {
		t.Fatalf(err.Error())
	}
	version = version >> 4
	if version != 1 {
		t.Errorf("invalid version expected 1 got %d -> %s", version, string(u))
	}
}

func TestVariant(t *testing.T) {
	u, err := Generate()
	if err != nil {
		t.Fatalf(err.Error())
	}
	variant, err := strconv.ParseUint(string(u[19]) + string(u[20]), 16, 64)
	if err != nil {
		t.Fatalf(err.Error())
	}
	variant = variant & 0xc0
	if variant != 0x80 {
		t.Errorf("invalid variant expected 8 got %d -> %s", variant, string(u))
	}
}

func TestCapacity(t *testing.T) {
	u, err := Generate()
	if err != nil {
		t.Fatalf(err.Error())
	}

	if cap(u) != 36 {
		t.Errorf("expected capacity 36 actual %d", cap(u))
	}
}


func BenchmarkGenerate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Generate()
	}
}

func BenchmarkGenerateParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Generate()
		}
	})
}

package gotest

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func sub1(t *testing.T) {
	t.Log("this is sub1")
}

func sub2(t *testing.T) {
	t.Log("this is sub2")
}
func sub3(t *testing.T) {
	t.Log("this is sub3")
	require.Equal(t, 1, 2)
}

func TestSub(t *testing.T) {
	t.Log("TestSub Begin")
	t.Run("sub1", sub1)
	t.Run("sub2", sub2)
	t.Run("sub3", sub3)
	// sub1(t)
	// sub2(t)
	// sub3(t)
	t.Log("TestSub End")
}

func BenchmarkSetBytes(b *testing.B) {
	b.Log("BenchmarkSetBytes Begin N=", b.N)
	b.ReportAllocs()
	b.SetBytes(1024 * 1024)
	for i := 0; i < b.N; i++ {
		sl := make([]byte, 1024*1024)
		_ = sl
		time.Sleep(100 * time.Millisecond)
	}
}

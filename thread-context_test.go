package threadContext

import (
	"github.com/stretchr/testify/assert"
	"runtime"
	"strconv"
	"sync"
	"testing"
)

const tCnt = 10000

func Test_GetMeta(t *testing.T) {
	memStats := new(runtime.MemStats)

	var wg sync.WaitGroup
	wg.Add(tCnt)
	for i := 0; i < tCnt; i++ {
		go func(want string) {
			defer wg.Done()

			SetMeta(nil, want)
			_, got := GetMeta()
			assert.Equal(t, want, got, "want", want, "got", got)
		}(strconv.Itoa(i))
	}
	wg.Wait()

	runtime.ReadMemStats(memStats)
	println("Mallocs:", memStats.Mallocs, "HeapObjects:", memStats.HeapObjects)

	runtime.GC()
	runtime.ReadMemStats(memStats)
	println("Frees:", memStats.Frees, "HeapObjects:", memStats.HeapObjects)
}

func BenchmarkSetMeta(b *testing.B) {
	var wg sync.WaitGroup
	wg.Add(b.N)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		go func(want string) {
			defer wg.Done()

			SetMeta(nil, want)
			GetMeta()
		}(strconv.Itoa(i))
	}
	wg.Wait()
}

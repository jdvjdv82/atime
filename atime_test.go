package atime

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/segmentio/encoding/json"
)

var testPointer *time.Time
var testAtime *AtomicTime

func TestZeroAllocJson(t *testing.T) {
	testAtime = New()
	bb, _ := json.Marshal(testAtime)
	if string(bb) != "null" {
		t.Fail()
	}
}

func TestNowAllocJson(t *testing.T) {
	dt := time.Now()
	testAtime = NewNow()
	bb, _ := json.Marshal(testAtime)
	if fmt.Sprintf("%q", dt.Format(time.RFC3339)) != string(bb) {
		t.Log("expect :", fmt.Sprintf("%q", dt.Format(time.RFC3339)))
		t.Log("got    :", string(bb))
		t.Fail()
	}

}

func TestTimeAllocJson(t *testing.T) {
	dt := time.Now()
	testAtime = NewTime(dt)
	bb, _ := json.Marshal(testAtime)
	if fmt.Sprintf("%q", dt.Format(time.RFC3339)) != string(bb) {
		t.Log("expect :", fmt.Sprintf("%q", dt.Format(time.RFC3339)))
		t.Log("got    :", string(bb))
		t.Fail()
	}

}

// read

func BenchmarkMutexTimeRead(b *testing.B) {
	mu := sync.RWMutex{}
	dt := time.Now().UTC()
	testPointer = &dt
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		mu.RLock()
		_ = testPointer
		mu.RUnlock()

	}
}

func BenchmarkATimeReadTime(b *testing.B) {
	testAtime = New()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = testAtime.GetTime()
	}
}

func BenchmarkATimeReadUnix(b *testing.B) {
	testAtime = New()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = testAtime.GetUnixTime()
	}
}

// write

func BenchmarkMutexTimePointerWrite(b *testing.B) {
	mu := sync.RWMutex{}
	dt := time.Now().UTC()
	testPointer = &dt

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		dt := time.Now().UTC()
		mu.Lock()
		testPointer = &dt
		_ = testPointer
		mu.Unlock()
	}
}

func BenchmarkMutexTimeWrite(b *testing.B) {
	mu := sync.RWMutex{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		mu.Lock()
		testTime := time.Now().UTC()
		_ = testTime
		mu.Unlock()
	}
}

func BenchmarkATimeWrite(b *testing.B) {
	testAtime = New()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testAtime.SetNow()
	}
}

// go test -bench=. -run=^Bench -v -benchtime 5s -benchmem

package atime

import (
	"sync/atomic"
	"time"

	"github.com/segmentio/encoding/json"
)

type AtomicTime struct {
	_ noCopy
	v int64
}

// noCopy may be embedded into structs which must not be copied
// after the first use.
//
// See https://github.com/golang/go/issues/8005#issuecomment-190753527
// for details.
type noCopy struct{}

// Lock is a no-op used by -copylocks checker from `go vet`.
func (*noCopy) Lock() {}

func New() *AtomicTime {
	return new(AtomicTime)
}

func NewNow() *AtomicTime {
	at := new(AtomicTime)
	at.SetNow()
	return at
}

func NewTime(t time.Time) *AtomicTime {
	at := new(AtomicTime)
	at.SetToTime(t)
	return at
}

func (at *AtomicTime) SetNow() {
	atomic.StoreInt64(&at.v, time.Now().UTC().Unix())
}

func (at *AtomicTime) SetNil() {
	atomic.StoreInt64(&at.v, 0)
}

func (at *AtomicTime) SetToTime(t time.Time) {
	atomic.StoreInt64(&at.v, t.Unix())
}

func (at *AtomicTime) GetUnixTime() int64 {
	return atomic.LoadInt64(&at.v)
}

func (at *AtomicTime) GetTime() time.Time {
	return time.Unix(atomic.LoadInt64(&at.v), 0)
}

func (at *AtomicTime) GetTimePointer() *time.Time {
	t := atomic.LoadInt64(&at.v)
	if t == 0 {
		return nil
	}
	tt := time.Unix(t, 0)
	return &tt

}

func (at *AtomicTime) SinceNow() time.Duration {
	return time.Since(time.Unix(atomic.LoadInt64(&at.v), 0))
}

// MarshalJSON behaves the same as if the AtomicTime is a *time.Time.
//
// I prefer for get null when time is not set over 1970-01-01...
func (at *AtomicTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(at.GetTimePointer())
}

// UnmarshalJSON behaves the same as if the AtomicTime is a *time.Time.
func (at *AtomicTime) UnmarshalJSON(b []byte) error {
	var v time.Time
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}

	at.SetToTime(v)
	return nil
}

package atime

import (
	"sync/atomic"
	"time"

	"github.com/segmentio/encoding/json"
)

type AtomicTime int64

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
	atomic.StoreInt64((*int64)(at), time.Now().UTC().Unix())
}

func (at *AtomicTime) SetNil() {
	atomic.StoreInt64((*int64)(at), 0)
}

func (at *AtomicTime) SetToTime(t time.Time) {
	atomic.StoreInt64((*int64)(at), t.Unix())
}

func (at *AtomicTime) GetUnixTime() int64 {
	return atomic.LoadInt64((*int64)(at))
}

func (at *AtomicTime) GetTime() time.Time {
	return time.Unix(atomic.LoadInt64((*int64)(at)), 0)
}

func (at *AtomicTime) GetTimePointer() *time.Time {
	t := atomic.LoadInt64((*int64)(at))
	if t == 0 {
		return nil
	}
	tt := time.Unix(t, 0)
	return &tt

}

func (at *AtomicTime) SinceNow() time.Duration {
	return time.Since(time.Unix(atomic.LoadInt64((*int64)(at)), 0))
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

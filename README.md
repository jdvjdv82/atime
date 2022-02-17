# ATime

Atomic Time package for Go, optimized for performance yet simple to use.

## Usage

```
// one line create
dt := atime.New() // allocates *AtomicTime
dt := atime.NewNow() // allocates *AtomicTime and set it to time.Now().UTC()
dt := atime.NewTime(time.Now()) // allocates *AtomicTime and set it your given time

// change values
dt.SetNow()  // set to time.Now().UTC()
dt.SetNil() // set to 0, expect to get null from json.Marshall
dt.SetToTime(time.Now()) // set it your given time

// get values
dt.GetUnixTime  // get unix time as int64
dt.GetTime() // get time as time.Time
dt.GetTimePointer() // get time as *time.Time

// get duration from Now
dt.SinceNow()

// embedding
type Foo struct {
  myTime *atime.AtomicTime // always use pointer to avoid copy
}
```

## Benchmark:

- Go 1.7.7
- Windows 11 64bit
- AMD Ryzen Threadripper 3960X 24-Core Processor

```
#read
BenchmarkMutexTimeRead-48               661886667                8.722 ns/op           0 B/op          0 allocs/op
BenchmarkATimeReadTime-48               1000000000               0.5250 ns/op          0 B/op          0 allocs/op
BenchmarkATimeReadUnix-48               1000000000               0.3965 ns/op          0 B/op          0 allocs/op

#write
BenchmarkMutexTimePointerWrite-48       134787126               44.44 ns/op           24 B/op          1 allocs/op
BenchmarkMutexTimeWrite-48              314054092               18.76 ns/op            0 B/op          0 allocs/op
BenchmarkATimeWrite-48                  867054152                6.804 ns/op           0 B/op          0 allocs/op
  
```


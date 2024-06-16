package pointer

import (
	"time"
)

// Bool returns a pointer to a bool.
func Bool(b bool) *bool {
	return &b
}

// Uint64 returns a pointer to a uint64.
func Uint64(u uint64) *uint64 {
	return &u
}

// Int returns a pointer to a int.
func Int(i int) *int {
	return &i
}

// Int32 returns a pointer to a int32.
func Int32(i int32) *int32 {
	return &i
}

// Int64 returns a pointer to a int64.
func Int64(i int64) *int64 {
	return &i
}

func Time(time time.Time) *time.Time {
	return &time
}

// To returns a pointer to the given value.
func To[T any](v T) *T {
	return &v
}

// Float64 returns a pointer to a float64.
var Float64 = To[float64]

// String returns a pointer to a string.
var String = To[string]

func ValueOf[T any](v *T) T {
	if v == nil {
		var zero T
		return zero
	}
	return *v
}

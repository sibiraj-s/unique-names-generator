package unique

import "time"

type basicType interface {
	bool | string | float64 | time.Time
}

// Ptr returns a pointer to the provided value.
func Ptr[T basicType](v T) *T {
	return &v
}

// String returns a pointer to the provided string.
func String(v string) *string { return Ptr(v) }

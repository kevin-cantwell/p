// Package p returns pointers of built in types. This is useful
// when it's not possible to take the address of an expression, such as
// a string literal or function call.
//
// The letter 'p' stands for pointer.
package p

// Given any value, return it's pointer
func Ptr(v interface{}) interface{} { return &v }

// Predeclared Identifiers

func Bool(b bool) *bool                   { return &b }
func Byte(b byte) *byte                   { return &b }
func Complex64(c complex64) *complex64    { return &c }
func Complex128(c complex128) *complex128 { return &c }
func Error(c error) *error                { return &c }
func Float32(f float32) *float32          { return &f }
func Float64(f float64) *float64          { return &f }
func Int(n int) *int                      { return &n }
func Int8(n int8) *int8                   { return &n }
func Int16(n int16) *int16                { return &n }
func Int32(n int32) *int32                { return &n }
func Int64(n int64) *int64                { return &n }
func Rune(r rune) *rune                   { return &r }
func String(s string) *string             { return &s }
func Uint(n uint) *uint                   { return &n }
func Uint8(n uint8) *uint8                { return &n }
func Uint16(n uint16) *uint16             { return &n }
func Uint32(n uint32) *uint32             { return &n }
func Uint64(n uint64) *uint64             { return &n }
func Uintptr(n uintptr) *uintptr          { return &n }

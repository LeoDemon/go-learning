package type_test

import (
	"math"
	"testing"
)

type MyInt int64

// go-lang does not support implicit type convert
func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
	t.Log(math.MaxInt8, math.MaxInt16, math.MaxInt32, math.MaxInt64)
	t.Log(math.MaxFloat32, math.MaxFloat64)
	t.Log(math.MaxUint8, math.MaxUint16, math.MaxUint32)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	// aPtr = aPtr + 1 // cannot using pointer for calculating
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("["+s+"]", len(s))
	if s == "" {
		t.Log("the default value of string is a empty string, not nil")
	}
}

package constant_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

const (
	Readable   = 1 << iota // 0001
	Writeable              // 0010
	Executable             // 0100
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)
	t.Log(Readable, Writeable, Executable)
	a := 7
	t.Log(a&Readable, a&Writeable, a&Executable)
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executable == Executable)
	a = 1
	t.Log(a&Readable, a&Writeable, a&Executable)
}

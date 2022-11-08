package main

import (
	"fmt"
)

type (
	Signed interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64
	}
	Unsigned interface {
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
	}
	Float interface {
		~float32 | ~float64
	}
	// --------------------------------------------------------------

	Integer interface {
		Signed | Unsigned
	}
	Numbers interface {
		Integer | Float
	}
	SomeInt interface {
		MyInt | Float | Integer
	}
)

type (
	MyInt   int
	MySlice []MyInt
)

func (ms MySlice) String() string {
	res := "{"
	for i, v := range ms {
		res += fmt.Sprintf("[%d: %v]", i, v)
	}
	return res + "}"
}

func (m MyInt) String() string {
	return fmt.Sprintf("%06d", m)
}

// Scale  returns a copy of s with each element multiplied by c.
func Scale[E Numbers](s []E, c E) []E {
	r := make([]E, len(s))
	for i, v := range s {
		r[i] = v * c
	}
	return r
}

func MakeMySlice[S ~[]T, T any](in ...T) S {
	r := make(S, len(in))
	for idx, v := range in {
		r[idx] = v
	}
	return r
}

// Scale2  returns a copy of s with each element multiplied by c.
func Scale2[S ~[]E, E Numbers](s S, c E) S {
	r := make(S, len(s))
	for i, v := range s {
		if v == E(42) {
			r[i] = v
			continue // do not touch holly digit!!! ;))))
		}
		r[i] = v * c
	}
	return r
}

func ScaleAndPrint(p MySlice, mul MyInt) {
	r := Scale(p, mul)
	// fmt.Println(r.String()) // DOES NOT COMPILE
	fmt.Println(r)
}

func ScaleAndPrint2(p MySlice, mul MyInt) {
	r := Scale2(p, mul)
	// fmt.Println(r.String())
	fmt.Println(r)
}

func main() {
	var m MyInt = 2
	var sl MySlice
	_ = m.String()
	m += 2
	ScaleAndPrint(MySlice{1, 2, 3, 4}, m)
	ScaleAndPrint2([]MyInt{1, 2, 3, 4}, 2)
	sl = make(MySlice, 0, 20)
	for i := 23; i <= 42; i++ { // 42 = answer for questions about "everything"
		sl = append(sl, MyInt(i))
	}
	fmt.Println(sl)
	ScaleAndPrint2(sl, m)
	sl2 := MakeMySlice[MySlice, MyInt](1, 2, 3, 4, 5, 6, 7, 7)
	sl3 := MakeMySlice[[]string, string]("1--", "--2", "++3", "4", "5", "6", "7, 7")
	fmt.Println(sl2, sl3)
}

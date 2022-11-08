package main

import (
	"errors"
	"fmt"
)

// Create a named type for our new error type.
type errorString string

// Implement the error interface.
func (e errorString) Error() string {
	return string(e)
}

// New creates interface values of type error.
func New(text string) error {
	return errorString(text)
}

var (
	ErrNamedType  = New("EOF")
	ErrStructType = errors.New("EOF")
)

func castError(err error) {
	switch err := err.(type) {
	case nil:
		fmt.Println("no error")
	case temporary:
		fmt.Printf("is temorary: %v\n", err.Temporary())
	case *MyError:
		fmt.Println("error occurred on line:", err.Line)
	default:
		fmt.Printf("unknown error: %s\n", err.Error())
	}
}

var someError = BadFunc()

func main() {
	if ErrNamedType == New("EOF") {
		fmt.Println("Named Type Error")
	}

	if ErrStructType == errors.New("EOF") {
		fmt.Println("Struct Type Error")
	}

	err := BadFunc()
	castError(err)
	if errors.Is(err, someError) {
		fmt.Println("this is someError")
	}
	var tmp *MyError
	if errors.As(err, &tmp) {
		fmt.Printf("got : %+v\n", *tmp)
	}

	st := []int{10, 12, 15}
	ist := make([]interface{}, len(st))
	for i := range st {
		ist[i] = st[i]
	}

	fmt.Println(DoIfaces(ist...))
}

type MyError struct {
	Msg  string
	File string
	Line int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%s:%d: %s", e.File, e.Line, e.Msg)
}

func (e *MyError) Temporary() bool {
	return false
}

func BadFunc() error {
	return &MyError{"Something happened", "server.go-lessons", 42}
}

type temporary interface {
	Temporary() bool
}

// IsTemporary returns true if err is temporary.
func IsTemporary(err error) (bool, error) {
	if err == nil {
		return false, errors.New("nil error")
	}
	te, ok := err.(temporary)
	return ok && te.Temporary(), nil
}

func DoIfaces(slice ...interface{}) error {
	return nil
}

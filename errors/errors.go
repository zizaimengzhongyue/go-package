package main

import (
	"errors"
	"fmt"
)

func Errors() {
	err := errors.New("test errors")
	fmt.Printf("%+v\n", err)
}

func Unwrap() {
	e := errors.New("test errors")
	err := fmt.Errorf("test wrap errors: %w", e)
	fmt.Printf("%+v\n", err)
	fmt.Printf("%+v\n", errors.Unwrap(err))
}

func Is() {
	e := errors.New("test errors")
	err := fmt.Errorf("test wrap errors: %w", e)
	fmt.Println(errors.Is(err, e))
	fmt.Println(errors.Is(e, err))
}

func As() {
	e := errors.New("test errors")
	err := fmt.Errorf("test wrap errors: %w", e)
	fmt.Println(errors.As(err, &e))
}

func main() {
	Errors()
	Unwrap()
	Is()
	As()
}

package main

import (
	"fmt"
	"github.com/goerr/goerr"
)

var (
	ErrGopher = fmt.Errorf("Gopher failed")
)

func pass0() (error, string, int) {
	return ErrGopher, "dog", 7
}

func pass1() (bool, error, int) {
	return true, ErrGopher, 9
}

func pass2() (bool, string, error) {
	return false, "cat", ErrGopher
}

func try() {
	errPASS0(pass0())
	errPASS1(pass1())
	errPASS2(pass2())

	fmt.Println("begin")

	u, _ := errPASS0(pass0())
	_, v := errPASS1(pass1())
	w, _ := errPASS2(pass2())

	m, n := errPASS0(pass0())
	p, q := errPASS1(pass1())
	x, y := errPASS2(pass2())

	fmt.Println(m, n, p, q, x, y, u, v, w)
}

func main() {
	goerr.OR0(try)
}

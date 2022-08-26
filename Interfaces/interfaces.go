package main

import "fmt"

type Test interface {
	Add() int
	Sub() int
}

type addSub struct {
	a int
	b int
}

func (c addSub) Sub() int {
	if c.a >= c.b {
		return c.a - c.b
	}
	return c.b - c.a
}

func (c addSub) Add() int {

	return c.a + c.b
}

func printVal(t Test) {
	fmt.Println(t.Add())
	fmt.Println(t.Sub())
}

func main() {
	a := addSub{
		a: 1, b: 1,
	}

	printVal(a)
}

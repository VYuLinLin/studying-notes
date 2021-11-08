package main

import (
	"fmt"
)

type Any interface{}
type EvalFunc func(Any) (Any, Any)

// 惰性求值 闭包函数（适合简单运算）
func main() {
	evenFunc := func(state Any) (Any, Any) {
		os := state.(int)
		ns := os + 2
		return os, ns
	}

	even := BuildLazyIntEvaluator(evenFunc, 0)

	for i := 0; i < 10; i++ {
		fmt.Printf("%vth even: %v\n", i, even())
	}
	// 0th even: 0
	// 1th even: 2
	// 2th even: 4
	// 3th even: 6
	// 4th even: 8
	// 5th even: 10
	// 6th even: 12
	// 7th even: 14
	// 8th even: 16
	// 9th even: 18
	fmt.Println("facttail", facttail(5, 6))
	fmt.Println("fn", fn(5, 6))
}
func fn(n, a int) int {
	if n == 1 {
		return a
	}
	return fn(n-1, a*n)
}
func facttail(n, a int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	if n == 1 {
		return a
	}
	return facttail(n-1, n*a)
}
func BuildLazyEvaluator(evalFunc EvalFunc, initState Any) func() Any {
	val := initState
	var newVal Any
	retFunc := func() Any {
		newVal, val = evalFunc(val)
		return newVal
	}
	return retFunc
}

func BuildLazyIntEvaluator(evalFunc EvalFunc, initState Any) func() int {
	ef := BuildLazyEvaluator(evalFunc, initState)
	return func() int {
		return ef().(int)
	}
}

package main

import "fmt"

import "rsc.io/quote"

func hello() {
    fmt.Println(quote.Go())

	var num int = 42
	fmt.Println("The number is:", num)

	var floatNum float64 = 3.14
	fmt.Println("The float number is:", floatNum)

	var isGoFun bool = true
	fmt.Println("Is Go fun?", isGoFun)


}
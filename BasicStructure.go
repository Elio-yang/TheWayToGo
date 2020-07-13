package main

import (
	"fmt"
)

//type IZ int

func main() {
	fmt.Println("hello world")
	func1()

	var a2 float32 = 0.45
	var b2 int = 7
	fmt.Println(Sum(a2, b2))

	var c int = 5
	fmt.Println(c)

	var d string = "hello elioyang"
	fmt.Println(d)

}
func func1() {
	fmt.Println("this is a function call")
}
func Sum(a float32, b int) int {
	return int(a) + b
}

/*
function:

	func functionName(para_list)(return_value_list){
		...
	}

	para_list (para1 type,para2 type,...)
	return_value_list (ret1 type,tet2 type,...)
*/

/*
type:

	declared by var will init with zero_value of this type
	basic: int float bool string
	complex: struct array slice map channel
	behaviour: interface

	for complex the default value is nil like null in java,there
	is no inheritance in Go_type.

	function can be a kind of return type:
		func funcName(para_list) typeFunc
	then somewhere in a function can return var with is of typeFunc:
		return var

	function with more than one return value:
		func funcName(para_list)(t1 type, t2 type){
			...
			return t1,t2
		}

	define our own type:
		type IZ int
		var a IZ=5
	this way is OK:
		type(
			IZ int
			FZ float64
			STR string

	)



*/

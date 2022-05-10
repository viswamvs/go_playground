package functions

import (
	"fmt"
	_"reflect"
)

func main(){
	answer := add(2,3)
	fmt.Println(answer)

	res, res1 := operation(10,5)
	fmt.Println(res, res1)

	factorial := fact(5)
	fmt.Println(factorial)

	result := calc(4,5, add)
	fmt.Println(result)

	output := addition(2,3)
	fmt.Println(output)

	//immediately invoked function expression
	sum := func (a,b int)int{
		return a + b
	}(2,3)
	
	fmt.Println(sum)
}

func add(a, b int) int{
	c := a+b
	return c
}

//named return values
func operation(a,b int)(add, mul int){
	add = a + b
	mul = a * b
	return
}

//recursive function
func fact(a int) int{
	if a <= 0 {
		return 1
	} else{
		return a * fact(a-1)
	}
}

//defer - defer is a keyword in Go that makes a function executes at the end of the execution of parent function

//function takes two arguments and third argument as a function
//derived type
type CalcFunc func (int, int) int
func calc(a int, b int, f CalcFunc) int{
	result := f(a,b)
	return result
}

//anonymous function (A function without a name)
var addition = func (a,b int) int {
	c := a + b
	return c
}

//pointer is a variable that points to memory location
//memory address are represented in hexadecimal values

package main

import "fmt"

/*
	A pointer also saves the memory address but it knows where that memory is located in the RAM 
	and how to retrieve the value stored in that memory address. 
	It can perform various kinds of operations with it like it can read the value stored at the memory address 
	or write a new value.

	'&' - referencing operator
	'*' - dereferencing operator

	Pointers have the power to mutate data they are pointing to
*/

func main() {
	a := 0x00

	//accessing memory address of the variable, add '&' in front of variable
	fmt.Println(&a)

	//(*int) means points to memory address of int data

	//create or declaring a pointer
	b := 1
	ptr := &b

	//dereferencing a pointer
	//fmt.Printf("ptr is a type of %T with value %v ", ptr, *ptr)
	
	//chaning variable value using a pointer
	*ptr = 23
	fmt.Println(b)

	//new function - allocates memory and returns pointer to that memory
	pt := new(int)
	fmt.Printf("pt is a %v with value %v", pt, *pt)
	fmt.Println()

	
	ans := 1
	
	//passing address of the variable
	doSomething(&ans)
	fmt.Println(ans)

	arr :=[3]int{1,2,3}
	simple(&arr)
	fmt.Println(arr)
}

//passing pointer to a function
func doSomething(ptr *int){
	*ptr = 100
}

//passing pointer of array type to a function
func simple(ptr *[3]int){
	(*ptr)[0] *= 2
	(*ptr)[1] *= 3
	(*ptr)[2] *= 4
}
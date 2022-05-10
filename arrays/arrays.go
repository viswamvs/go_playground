package arrays

import "fmt"

func main(){

	//method 1 - to create a array
	var arr [5]int
	arr[0] = 100
	arr[1] = 200
	fmt.Println(arr)

	//method 2 - shorthand way to create array
	arr1 := [3]string {"viswa", "nikhil", "rohan"}
	fmt.Println(arr1)

	//method 3 - initializing array with an ellipses (Automatic array length declaration)
	arr2 := [...]int{1,2,3,4,5}
	fmt.Println(len(arr2))

	//iteration of array
	for index, value := range arr2{
		fmt.Println(index, "->", value)
	}
}
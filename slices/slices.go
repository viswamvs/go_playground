//slice is just reference to a array
package slices

import "fmt"

func main() {

	//creation of slice using make function (empty slice)
	s := make([]int, 3, 10)
	s[0] = 1
	s[1] = 2
	s[2] = 3
	fmt.Println(cap(s))

	arr := [10]int{1,2,3,4,5,6,7,8,9,10}
	sl := arr[2:4]
	fmt.Println("New Slice ", sl)
	fmt.Println("Capacity ", cap(sl))

	//creating another slice from array
	sl1 := arr[2:4]

	//appending slice values to another one
	newSlice := append(sl1, 100,200)
	fmt.Println(arr)
	fmt.Println("Slice 2", sl1, newSlice)

	//address of second element in an array
	fmt.Println("Address of array", &arr[2])

	//address of first element in slice 
	//both arr[2] and sl[0] points to same element, so address will be same
	fmt.Println("Address of slice", &sl[0])

	//changing the value in slice affects the array also
	sl[0] = 475
	fmt.Println(arr[2])
	
	//append function
	s = append(s, 4)
	s = append(s, 5)
	fmt.Println(s)

	s1 := s[1:]
	fmt.Println("Slice s1",s1) // 2, 3, 4, 5

	s2 := s[0:3]
	fmt.Println("Slice s2",s2) // 1, 2, 3

	s3 := s[:3]
	fmt.Println("Slice s3",s3) // 1, 2, 3

	//another way of creating slices
	names := []string{"viswa", "nikhil", "rohan"}
	fmt.Println(names)

	//copy function (copying slice into another slice)
	//copy doesn't append new elements, it replaces only the existing elements
	var slice1 []int //nil slice
	var slice2 []int = []int{1,2,3,400}
	var slice3 []int = []int{4,5,6}
	var slice4 []int = []int{1,2,3}

	n1 := copy(slice1, slice2)
	//slice1 is nil slice, nothing will happen and slice1 will be nil
	fmt.Println(n1, slice1) // 0, [] (empty slice)

	n2 := copy(slice2, slice3)
	fmt.Println(n2,slice2) //

	n3 := copy(slice4, slice3)
	fmt.Println(n3,slice4)

	var arrSlice []int = []int{0,1,2,3,4,5}
	//here slice is passed as reference, it affects the arrSlice
	doSomething(arrSlice)

	fmt.Println(arrSlice)

	var zeroValueSlice []int
	fmt.Println(zeroValueSlice == nil)

	arrays := [10]int{10,20,30,40,50,60,70,80,90,100}

	//capacity of slice 
	//In below example, slice references the array from index 2. Hence the capacity of slice
	//will be 8 (because index 2 to index 9) 
	sliceArray := arrays[8:9]
	fmt.Println(sliceArray)
	fmt.Println("Capacity of New Slice ", cap(sliceArray))

	array := [9]int{1,2,3,4,5,6,7,8,9}

	//slice contains array elements from index 2 to index 4 (i.e., 3 and 4)
	arrS := array[2:4]

	//newS has elements from arrS and 55 & 66
	newS := append(arrS, 55, 66)

	fmt.Println(arrS, newS)

	//change in newS also affects the original array that references
	fmt.Println(array) // output : [1,2,3,4,55,66,7,8,9]

	//unpack operator - In slice, to append values to it from another slice
	unpackSlice := make([]int, 0, 10)
	sliceUnpack := []int{1,2,3}

	sliceUnpack = append(sliceUnpack, unpackSlice...)
	fmt.Println("Appending values from another slice",sliceUnpack)

}

//making squares from slice
func doSomething(slice []int) {
	for index, value := range slice {
		slice[index] = value * value
	}
}
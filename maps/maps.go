package maps

import "fmt"

func main() {

	//creating a map with make function
	m := make(map[int]string)
	m[1] = "viswa"
	m[2] = "nikhil"
	m[3] = "rohan"

	//fmt.Println(len(m))

	//delete a key-value pair in map
	delete(m, 3)
	//fmt.Println(m)

	//another way of creating map
	m1 := map[string]int{"viswa":123}
	fmt.Println(m1)

	var myMap map[int]string
	myMap = m
	fmt.Println(myMap)

	age := map[string]int{
		"viswa":22,
	}
	_,b := age["viswa"]
	c, d:= age["viswaaa"] 
	//fmt.Println(age["viswa"])
	fmt.Println(b)
	fmt.Println(c,d)

	delete(age, "viswa")
	fmt.Println(age)

	newMap := make(map[int]string)

	newMap[0] = "a"
	newMap[1] = "b"
	newMap[2] = "c"
	newMap[3] = "d"

	//iterating map using range
	for key, value := range newMap {
		fmt.Println(key , "=>" , value)
	}
}
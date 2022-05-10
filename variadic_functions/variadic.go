package variadicfunctions

import "fmt"

func main(){
	nums := []int{1,2,3,4,5}
	totalSum(nums...)

	totalSum(1,2,3)
	totalSum(0,15,20)
	totalSum(1,2)
}

//variadic functions - can be called with any number of trailing arguments
func totalSum(nums ...int){
	total := 0

	for _, value := range nums{
		total += value
	}

	fmt.Println(total)
}
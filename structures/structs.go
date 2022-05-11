package structures

import (
	"fmt"
)

//Declaring struct type
type Salary struct {
	basic int
	insurance int
	allowance int
}
type Employee struct {
	firstName, lastName string
	Salary //nested struct
	fullTime bool
}


func main(){
	//creating struct (here employee is a struct)
	var employee Employee
	
	//setting values to struct fields (initializing struct)
	employee.firstName = "viswa"
	employee.lastName = "m"
	employee.Salary = Salary{1000, 200, 300}
	employee.fullTime = true

	//getting values from struct fields
	fmt.Println(employee.firstName)
	fmt.Println(employee.lastName)

	//one way of initializing struct
	//here we can change the order of fields and also we can initialize only some fields
	emp := Employee{
		firstName: "arun",
		lastName: "kumar",
		Salary: Salary{1200, 100, 2000},
		fullTime: false,
	}
	fmt.Println(emp)

	//another way of initializing struct
	//here it's mandatory to initialize all fields
	em := Employee{"dinesh", "ravi", Salary{1000, 200, 100}, false}
	fmt.Println(em)
	fmt.Printf("%T", em)
	fmt.Println()

	//creating anonymous struct
	employee1 := struct{
		firstName, lastName string
		salary int
	}{
		firstName: "nikhil",
		lastName: "saji",
	}
	fmt.Println(employee1)
	fmt.Printf("%T", employee1)
	fmt.Println()

	//creating pointer to a struct
	empl := &Employee{
		firstName: "viswa",
		lastName: "m",
		fullTime: true,
		Salary: Salary{2000, 0, 200},
	}

	//dereferencing (*empl) to get actual value
	fmt.Println((*empl).firstName)

	//we can access the struct poiner without dereferencing it
	fmt.Println(empl.fullTime)

	//struct type of anonymous fields
	type Data struct{
		string
		int
		bool
	} 

	sample := Data{"viswa", 1000, false}
	fmt.Println(sample.bool)

	//field promotion - anonymous nested struct, all nested struct fields are automatically 
	//available on parent struct
}
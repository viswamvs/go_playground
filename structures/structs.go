package main

import "fmt"

//Declaring struct type
type Employee struct {
	firstName, lastName string
	salary int
	fullTime bool
}

func main(){
	//creating struct (here employee is a struct)
	var employee Employee
	
	//setting values to struct fields (initializing struct)
	employee.firstName = "viswa"
	employee.lastName = "m"
	employee.salary = 1200
	employee.fullTime = true

	//getting values from struct fields
	fmt.Println(employee.firstName)
	fmt.Println(employee.lastName)

	//one way of initializing struct
	//here we can change the order of fields and also we can initialize only some fields
	emp := Employee{
		firstName: "arun",
		lastName: "kumar",
		salary: 1000,
		fullTime: false,
	}
	fmt.Println(emp)

	//another way of initializing struct
	//here it's mandatory to initialize all fields
	em := Employee{"dinesh", "ravi", 100, false}
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
		salary: 1000,
	}

	//dereferencing (*empl) to get actual value
	fmt.Println((*empl).firstName)

	//we can access the struct poiner without dereferencing it
	fmt.Println(empl.fullTime)
}
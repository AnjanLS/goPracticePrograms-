// package main 
// import "fmt"

// type person struct{
// 	name string
// 	age int
// 	job string
// 	salary int
// }

// func main(){
// 	var person1 person
// 	var person2 person

// 	person1.name="Anjan"
// 	person1.age=24
// 	person1.job="software engineer"
// 	person1.salary=25000

// 	person2.name="Arjun"
// 	person2.age=24
// 	person2.job="Army"
// 	person2.salary=40000

// 	fmt.Printf("Name:%v\n",person1.name)
// 	fmt.Printf("Age:%v\n",person1.age)
// 	fmt.Printf("Job:%v\n",person1.job)
// 	fmt.Printf("Salary:%v\n",person1.salary)

// 	fmt.Printf("Name:%v\n",person2.name)
// 	fmt.Printf("Age:%v\n",person2.age)
// 	fmt.Printf("Job:%v\n",person2.job)
// 	fmt.Printf("Salary:%v\n",person2.salary)

// }


package main 
import "fmt"

type person struct{
	name string
	age int
	job string
	salary int
}

func main(){
	var person1 person
	var person2 person

	person1.name="Anjan"
	person1.age=24
	person1.job="software engineer"
	person1.salary=25000

	person2.name="Arjun"
	person2.age=24
	person2.job="Army"
	person2.salary=40000

	printperson(person1)
	printperson(person2)
}

func printperson(person person){
	fmt.Printf("Name:%v\n",person.name)
	fmt.Printf("Age:%v\n",person.age)
	fmt.Printf("Job:%v\n",person.job)
	fmt.Printf("Salary:%v\n",person.salary)
}

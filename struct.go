package main
import "fmt"

type Customer struct{
	Name, Address string
	Age int
}

type Resiler struct{
	Name, Address string
	Age int
}

func (customer Customer) fanc_metohd(name string){
	fmt.Println("Hello", name, "myname is ", customer.Name)
}

func main(){
	var data Customer
	data.Name = "Andini"
	data.Address = "data Address"
	data.Age = 23
	
	// funct method
	data.fanc_metohd("Jhone")

	// struct literal
	// resiler := Resiler{
	// 	Name: "Joko",
	// 	Address: "kaliurang",
	// 	Age:23,
	// }


	// fmt.Println(data.Name)
	// fmt.Println(data.Address)
	// fmt.Println(data.Age)
	// fmt.Println(resiler)
}


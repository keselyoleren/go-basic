package main
import "fmt"

type HasName interface {
	GetName() string
}

func interface_func(hasname  HasName){
	fmt.Println("hello", hasname.GetName())
}

type Person struct {
	Name string
}

type Animal struct{
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func (person Person) GetName() string {
	return person.Name
}



// interface kosong
func Interface_kosong(param interface{}) interface{} {
	return param
}

func main(){
	data := Person{
		Name: "della",
	}
	interface_func(data)

	data_animal := Animal{
		Name: "Kucing",
	}
	interface_func(data_animal)

	var data_interface interface{} = Interface_kosong("strong")
	fmt.Println(data_interface)
}
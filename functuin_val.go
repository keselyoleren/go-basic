package main
import "fmt"

func ExampleFunc(name string) string{
	return name
}

func main(){
	var_func := ExampleFunc
	fmt.Println(var_func("test function"))
}
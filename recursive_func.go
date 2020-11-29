package main
import "fmt"

func factorialLoop(val int)int{
	result := 1
	for a:=val; a>0; a--{
		result *=a
	}
	return result
}

func recursive_func(value int) int{
	if value ==1 {
		return 1
	} else {
		return value * recursive_func(value-1)
	}

}

func main(){
	fac := recursive_func(5)
	fmt.Println(fac)
}
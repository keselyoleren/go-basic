package main
import "fmt"

func Variadic(numers ...int)int{
	total := 0
	for _, val:= range numers {
		total += val
	}
	return total
}

func main(){
	slice := []int{1,2,3,4,5,6,7}
	// variadic := Variadic(1,2,3,4,5)
	variadic := Variadic(slice...)
	fmt.Println(variadic)
}
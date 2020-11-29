package main
import "fmt"

func random() interface {} {
	return 1
}

func main(){
	var result interface{} = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	switch value := result.(type){
	case string:
		fmt.Println( value, "is string")
	case int:
		fmt.Println( value,  "is int")
	case bool:
		fmt.Println( value,  "is bools")
	default:
		fmt.Println( value ,"unconwon type")
	}
}
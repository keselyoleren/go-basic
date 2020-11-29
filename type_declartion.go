package main
import "fmt"

func main(){
	// alias variable
	type str_alias string

	var var_als str_alias = "alias declaration"
	fmt.Println(var_als)
}
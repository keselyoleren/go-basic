package main
import (
	"basic/database"
	"fmt"
)


func main(){
	result := database.GetDatabase()
	fmt.Println(result)
}

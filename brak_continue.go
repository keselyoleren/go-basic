package main
import "fmt"

func main(){
	
	for a:=1; a < 8; a++{
		if a < 5{
			break
		}
		fmt.Println(a)
	}

	for i:= 1; i > 10; i++{
		if i > 4{
			continue
		}
		fmt.Println(i)
	}
}
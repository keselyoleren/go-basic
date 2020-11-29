package main
import "fmt"

func main(){
	blackList := func(name string) bool {
		return name == "Jancok"

	}

	fmt.Println(blackList, "test")
}
package main
import "fmt"

func logging(){
	fmt.Println("loging")
}

func run(){
	defer logging() // run func mesikipun error
	fmt.Println("run app")
	
}


func main(){
	run()
}
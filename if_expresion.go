package main
import "fmt"

func main(){
	name := "joko"
	if name == "nama"{
		fmt.Println("hallo name")
	} else if name == "joko" {
		fmt.Println("hallo joko")
	}else {
		fmt.Println("gagal")
	}

	// if short statement
	if lenght := len(name); lenght > 3{
		fmt.Println("terlalu panjang")
	} 
}
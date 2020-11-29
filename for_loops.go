package main
import "fmt"

func main(){
	conter := 1
	
	for conter < 10 {
		fmt.Println(conter)
		conter++
	}

	// for demgam statement
	for count := 1; count < 10; count++ {
		fmt.Println(count)
	}

	slice := []string{"nama", "alamat", "kelamin"}

	for a := 0; a < len(slice); a++{
		fmt.Println(slice[a])
	}

	// for range
	for key, val := range slice{
		fmt.Println(key, val)
	}
}
package main
import "fmt"


// functio di dalam func haris hati2 karea bisa merubah data di dlm func tsb
func main(){
	counter := 0

	increment := func() {
		counter++
	}
	increment()

	fmt.Println(counter)
}
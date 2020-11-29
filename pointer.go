package main
import "fmt"

type Address struct {
	City, Province, Country string
}

func main(){
	address1 := Address{"Banyuwangi","Jawa timur", "Indonesia"}
	// pass by value
	// address2 := address1

	// pass by referebce
	address2 := &address1
	address2.City = "Makasar"
	fmt.Println(address2)
	fmt.Println(address1)

}
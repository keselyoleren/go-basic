package main
import "fmt"

func main(){
	name := "Dila"
	switch name {
	case "Dila":
		fmt.Println("hai Dilla")
	case "Novi":
		fmt.Println("hai Novi")
	default:
		fmt.Println("SUdah bahagiakah Anda")
	}

	// swith short statement
	switch length := len(name); length > 2 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("Nama anada kurang bagus")
	}

	// swhitch tanpa kondisi
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama kurang panjang")
	case length <5:
		fmt.Println("nama sudah benar")
	}
}
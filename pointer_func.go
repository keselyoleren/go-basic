package main
import "fmt"

type Adress struct {
	City, Province, Country string
}

func ChangeAddresstoInd(address *Adress){
	address.Country = "indonesia"
}


// pinter in method


func main(){
	alamat := Adress{
		City: "Banyuwangi",
		Province:"Jawa timur",
		Country: "Ind",
	}
	ChangeAddresstoInd(&alamat)
	fmt.Println(alamat)
}
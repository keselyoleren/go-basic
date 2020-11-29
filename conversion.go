package main
import "fmt"

func main(){
	// conversi type data
	// int32 to ont64
	var nilai32 int32 = 32768
	var nilai64 = int64(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)

	// conversi byte  to string
	name:= "string name"
	byte:= name[0]
	convert_to_strong := string(byte)
	fmt.Println(convert_to_strong)
}
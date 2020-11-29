package main
import "fmt"

func main(){
	var str string 
	str = "string"
	fmt.Println(str)
	str = "change value var"
	fmt.Println(str)

	// int
	var fr_int int8
	fr_int = 19
	fmt.Println(fr_int)

	fr_int = 20
	fmt.Println(fr_int)

	// kata var tidak wajib bisa diganti dengan :=
	var1 := "string"
	fmt.Println(var1)

	var1 = "change string"
	fmt.Println(var1)

	// multiple variable 
	var (
		nama = "test nama"
		alamat = "test alamat"
	)

	fmt.Println(nama)
	fmt.Println(alamat)

	// constatn variable
	const const_var string = "constant variable"
	fmt.Println(const_var)

	// const multiple var
	const (
		nama_const  = "const var nama"
		alamat_const  = "const var alamat"
	)

	fmt.Println(nama_const)
	fmt.Println(alamat_const)


	
	
}
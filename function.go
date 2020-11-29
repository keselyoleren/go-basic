package main
import "fmt"


func test(nama string){
	fmt.Println(nama)
}


// function resturn val
func Return_val(test string) string {
	return test
}

// return multiple val
func return_multiple_val() (nma string, alamat string){
	nma = "test"
	alamat = "test"
	return
}

func main(){
	fun := Return_val("test")
	test("hai saya")
	return_multiple_val()
	fmt.Println(fun)
}
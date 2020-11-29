package main
import "fmt"

type Man struct{
	Name string
}

func (man Man) Merried() {
	man.Name= "Hai, " + man.Name
}

func main(){
	data := Man{"Supri"}
	data.Merried()
	fmt.Println(data.Name)
}
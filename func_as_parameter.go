package main
import "fmt"

type Filter func(string) string
func FileterKata(name string, filter Filter){
	name_filtered := filter(name)
	fmt.Println("hello", name_filtered)	
}

func spamFilter(name string) string{
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}

} 

// with type declaration func


func main(){
	filter_spm := spamFilter("Anjing")
	fmt.Println(filter_spm)
}
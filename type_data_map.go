package main
import "fmt"

func main(){
	maps := map[string]string{
		"key1":"val1",
		"key2":"val2",
	}

	delete(maps, "key1")

	fmt.Println(len(maps))
	fmt.Println(maps["key1"])

}
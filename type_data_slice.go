package main
import "fmt"

func main(){
	var bulan = [...]string{
		"Januari",
		"februari",
		"Maret",
		"Aprel",
		"Mei",
		"Jumi",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"NOvember",
		"Desember",
	}
	var slice1 = bulan[4:5]

	fmt.Println(slice1)
	fmt.Println(cap(slice1))
	fmt.Println(len(slice1))
	fmt.Println(append(slice1, "thn Depan"))
}
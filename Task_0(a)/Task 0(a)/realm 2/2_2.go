package main
import "fmt"
func main(){
	var day int = 3
	switch{
	case day==1:
		fmt.Print("2 labs")
	case day==2:
		fmt.Print("only theory")
	case day==3:
		fmt.Print("who comes to college on this dayy")
	default:
		fmt.Print("bhaii college aajaa")

	}
}
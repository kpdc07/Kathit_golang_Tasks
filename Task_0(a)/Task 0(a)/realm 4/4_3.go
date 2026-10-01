package main
import "fmt"
func main(){
	type student struct {
		name string
		age int
		marks int 
	}
	st1 :=student{
		"kathit",18,15,
	}
	fmt.Println(st1)
}
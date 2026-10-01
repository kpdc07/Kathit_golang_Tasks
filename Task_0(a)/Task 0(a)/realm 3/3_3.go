package main
import "fmt"
func add( no ...int)(result int){
	total:=0
	for _,no := range no{
		total = total +no
	}
	return total
}
func main(){
	result := add(1,2,3,4,5,6)
	fmt.Println(result)
}
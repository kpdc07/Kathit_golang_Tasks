package main
import "fmt"
func calc ( a int, b int)( sum int, diff int){
	return a+b,a-b
}
func main(){
	a:=10
	b:=3
	sum,diff := calc(a,b)
	fmt.Println(sum)
	fmt.Println(diff)
}
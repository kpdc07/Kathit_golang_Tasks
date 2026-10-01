package main
import "fmt"
func add (a int,b int)(result int){
	result = a+b
	return
}
func main(){
a:=10
b:=20
result := add(a,b)
fmt.Println(result)
}
package main
import "fmt"
func main(){
	a :=[3]int{1,2,3}
	b :=[]int{1,2,3}
	b = append(b,4)
	for _,a := range a{
		fmt.Print(a)
	}
	for _,b := range b{
		fmt.Print(b)
	}
}
// array-> fixed length ,no changes after declaring
// sclice-> var length , changes can be made afterwards
package main
import "fmt"
func main(){
	marks :=[]int{10,20,30,40}
	for _,marks := range marks{
		fmt.Println(marks)
	}
}
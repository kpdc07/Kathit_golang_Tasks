package main
import "fmt"
func main(){
	for i:=0;i<3;i++{
		fmt.Println(i)
	}
	i :=3
	for i<6{
		fmt.Println(i)
		i++
	}
	for {
		fmt.Println("infinite untill breaked")
		break
	}
}
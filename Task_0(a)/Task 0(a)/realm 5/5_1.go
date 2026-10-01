package main
import "fmt"
import"time"
func hi(){
	fmt.Println("hello")
}
func main(){
	go hi()
	time.Sleep(time.Second)
	fmt.Println("hello main")
}
package main
import "fmt"
import "time"
func send(ch chan int){
	ch <-10
}
func receive(ch chan int){
	channel:= <-ch 
	fmt.Println("recvd.",channel)
}
func main(){
	ch:= make(chan int)
	go send(ch)
	go receive(ch)
	time.Sleep(time.Second)
}
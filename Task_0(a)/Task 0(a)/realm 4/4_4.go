package main
import "fmt"
type student struct{
	name string
	age int 
}
func (s *student)birthday(){
	s.age ++
	return 
}
func main(){
	s1 :=student{
		name : "kathit",
		age :18,
	}
	s1.birthday()
	fmt.Println(s1)
}
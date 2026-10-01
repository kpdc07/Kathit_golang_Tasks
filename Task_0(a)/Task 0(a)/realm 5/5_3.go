package main
import "fmt"
func div(a int, b int)(int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cant divide by zero")
	}
	return a/b ,nil
}
func main() {
	result,err := div(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
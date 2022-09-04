// operators
// 1. Comparison Operators
/*package main
import "fmt"

func main(){
	var city string = "seoul"
	var city2 string = "tokyo"
	fmt.Print(city == city2)
}*/
// 2. Arithmetic Operators

// 3. Logical Operators
// 3-1. &&
/*package main
import "fmt"

func main(){
	var x int = 10
	fmt.Println((x < 100) && (x < 200))
	fmt.Println((x < 300) && (x < 0))
}*/
//3-2. ||
package main
import "fmt"

func main(){
	var x int = 10
	fmt.Println((x < 0) || (x < 200))
	fmt.Println((x < 0) || (x > 200))
}
// 4. Assignment Operators

// 5. Bitwise Operators
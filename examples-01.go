//01.single input examples
/*package main
import "fmt"

func main() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &name)
	fmt.Println("Hey there, ", name)
}
*/

//02.multiple input
package main
import "fmt"

func main() {
	var name string
	var is_muggle bool

	fmt.Print("Enter your name & are you a muggle: ")
	fmt.Scanf("%s %t",&name, &is_muggle)
	fmt.Println(name, is_muggle)
}

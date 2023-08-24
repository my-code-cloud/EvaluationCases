//revive:disable
package fixtures

import "fmt"

//revive:enable:range-val-in-closure
func foo() {
	mySlice := []string{"A", "B", "C"}
	for index, value := range mySlice {
		go fmt.Printf("Index: %d\n", index) // MATCH /loop variable index captured by func literal/
		go fmt.Printf("Value: %s\n", value) // MATCH /loop variable value captured by func literal/
	}
}

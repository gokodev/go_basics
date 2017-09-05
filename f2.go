package main
import (
"fmt"
"os")
func main() {
//Alternative definition of variable type INT.
//var n int = 5
n := 5
if len(os.Args) !=n {
os.Exit(1)
}
fmt.Println("over", os.Args[1:n])
}
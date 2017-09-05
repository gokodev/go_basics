//In this program 'n' denotes the number of arguments that is considered to execute the IF condition.
//Theorotically if the number of arguments passed to the program is not equal to 5 the program should exit, else (equal to 5) then it must format and print.
//But practically, the IF condition is executed in a (n-1) fashion, so in this case if the number/length of arguments equals 4 then "fmt.Println..." is executed which is undesirable and needs investigation.
//Link: https://stackoverflow.com/questions/18973335/golang-check-number-of-arguments-also-user-input-check-for-return-key-blan
//Solution: https://stackoverflow.com/questions/30302017/golang-command-arguments-empty-causing-error
//By default, the first argument is the program itself (f2.go), hence you need to increment you required arguments by 1. Simple :)

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
package main
import (
 "fmt"
 "log"
)
func main() {
 var num1 = 5
 fmt.Println("Hello, world!")
}

//The presence of an unused variable may indicate a bug, while unused imports just 
//slow down compilation. Accumulate enough unused imports in your code tree and 
//things can get very slow. For these reasons, Go allows neither


//If you insist on having unused variables in your program, Go offers a quick way 
//to address this issue. Simply assign the unused variable to a blank identifier (_)
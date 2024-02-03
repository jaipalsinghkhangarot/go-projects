package main
import "fmt"
var num = 35 // type inferred
func main() {
 var num1 = 5 // type inferred
 var num2 int // explicitly typed
 var num3 float32 // floating point variable
 var name string // string variable
 var raining bool // boolean variable
 var rates float32 = 4.5 // declared as float32 and then initialized

 fName := "Wei-Meng-tee"
 firstName, lastName, age := "Wei-Meng", "Lee", 25
 var first_Name, last_Name string = "Weitee-Meng", "Lee"

 var (
	ftName string = "Wei-Mengchesd"
	ltName string = "Leegty"
	ageee int = 25
   )

 fmt.Println(num1) // 5
 fmt.Println(num2) //Variables declared without initialization are given their zero value
 fmt.Println(num3) // 0
 fmt.Println(name) // "" (empty string)
 fmt.Println(raining) // false
 fmt.Println(rates)
 fmt.Println(num) // 35

 fmt.Println(fName) 
 fmt.Println(firstName)
 fmt.Println(lastName) 
 fmt.Println(age) 

 fmt.Println(first_Name)
 fmt.Println(last_Name)

 fmt.Println(ftName)
 fmt.Println(ltName) 
 fmt.Println(ageee)
}
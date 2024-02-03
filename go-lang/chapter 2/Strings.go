package main
import(
	"fmt"
	"unicode/utf8"
)
func main()  {
var firstName string = "Wei-Meng"
address := "The White House\n1600 Pennsylvania Avenue NW\nWashington, DC 20500"
quotation := `"Anyone who has never made
a mistake has never tried
anything new."
--Albert Einstein`
str1 := "你好,世界" // Chinese
str2 := "こんにちは世界" // Japanese
str3 := "\u4f60\u597d\uff0c\u4e16\u754c" // 你好,世界
fmt.Println(firstName)
fmt.Println(address)
fmt.Println(quotation)
fmt.Println(len(str1)) // 15 = 5 chars * 3 bytes
fmt.Println(len(str2)) // 21 = 7 chars * 3 bytes
fmt.Println(len(str3)) // 15 = 5 chars * 3 bytes
fmt.Println(utf8.RuneCountInString(str1)) // 5
fmt.Println(utf8.RuneCountInString(str2)) // 7
fmt.Println(utf8.RuneCountInString(str3)) // 5
}

//If you want to count the number of characters (runes) in a string, use the 
//RuneCountInString() function:

//A rune is any of the characters of certain ancient alphabets (for example, a script 
	//used for writing the Germanic languages)
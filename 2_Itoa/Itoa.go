package Itoa

// import "fmt"
// func main() {
// 	var num int
// 	fmt.Scan(&num)
// 	s := Itoa(num)
// 	fmt.Printf("Result:%s, %T \n", s, s)
// }

func Itoa(num int) string {
	var str string
	if (num / 10) != 0 {
		str = Itoa(num / 10)
	}
	str += string('0' + (num % 10))
	return str
}

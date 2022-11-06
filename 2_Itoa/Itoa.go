package Itoa

// package main

// import "fmt"

// func main() {
// 	var num int
// 	fmt.Scan(&num)
// 	s := Itoa(num)
// 	fmt.Printf("Result:%s, %T \n", s, s)
// }

func Itoa(num int) string {
	var len int = 1
	var tmp int

	tmp = num
	for tmp != 0 {
		tmp = tmp / 10
		len++
	}

	var negBool bool
	var strByte = make([]byte, len)

	if num < 0 {
		negBool = true
		num *= -1
	}

	for i := len - 1; (num / 10) != 0; i-- {
		strByte[i] = byte('0' + num%10)
		num = num / 10
	}
	strByte[0] = byte('0' + num%10)

	if negBool {
		return "-" + string(strByte)
	} else {
		return string(strByte)
	}
}

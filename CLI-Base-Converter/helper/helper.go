package helper

import (
	"fmt"
	"strconv"
	"strings"
)

func HexConv(s string) int64 {
	num, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		fmt.Println("Error in conversion, Please enter a valid (Hex) value")
	}
	return num
}

func BinConv(s string) int64 {
	num, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		fmt.Println("Conversion Voided, Please input a valid Bin value")
	}
	fmt.Println("result: ")
	return num
}

func DecToHex(s string) string {
	c, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println("conversion unsuccesful")
		fmt.Println("Please Enter a valid decimal value")
	}
	d := strconv.FormatInt(int64(c), 16)
	fmt.Print("The HEX to BIN value is: ")

	return strings.ToUpper(d)
}

func DecToBin(s string) string {
	c, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println("conversion unsuccesful")
		fmt.Println("Please Enter a valid decimal value")
	}
	d := strconv.FormatInt(int64(c), 2)
	fmt.Print("The Dec to BIN value is: ")

	return strings.ToUpper(d)

}

// func DecToHex(s string) {

// }

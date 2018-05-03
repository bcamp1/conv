package main

/* Branson Camp
Arg 1:
-d	Decimal
-b	Binary
-h	Hex
-a	Ascii

Arg 2:
-d	Decimal
-b	Binary
-h	Hex
--base	Specific Base
-a
*/

import (
	"fmt"
	"os"
	"strconv"
)

func decToBase(dec, base int) string {
	return strconv.FormatInt(int64(dec), base)
}

func convertDec(dec int) string {
	type2 := os.Args[2]

	switch type2 {
	case "-d":
		return decToBase(dec, 10)
		break
	case "-b":
		return decToBase(dec, 2)
		break
	case "-h":
		return decToBase(dec, 16)
		break
	case "-a":
		return string(dec)
	default:
		if string(type2[0:6]) == "--base" {
			base, _ := strconv.Atoi(type2[6:])
			return decToBase(dec, base)
		}
	}
	return ""

}

func main() {
	type1 := os.Args[1]
	value := os.Args[3]
	var dec int
	var output string

	switch type1 {
	case "-d":
		dec, _ = strconv.Atoi(value)
		output = convertDec(dec)
		break
	case "-b":
		conversion, _ := strconv.ParseInt(value, 2, 64)
		dec = int(conversion)
		output = convertDec(dec)
		break
	case "-h":
		conversion, _ := strconv.ParseInt(value, 16, 64)
		dec = int(conversion)
		output = convertDec(dec)
		break
	case "-a":
		dec = int(value[0])
		output = convertDec(dec)
		break
	default:
		if string(type1[0:6]) == "--base" {
			base, _ := strconv.Atoi(type1[6:])
			conversion, _ := strconv.ParseInt(value, base, 64)
			dec = int(conversion)
			output = convertDec(dec)
		}
		break
	}

	fmt.Println(output)
}

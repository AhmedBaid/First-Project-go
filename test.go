package main

import "fmt"

func HexToByte(hex string) byte {
	var result byte
	for _, char := range hex {
				var value byte
		if char >= 'a' && char <= 'f' {
			char -= 32 
		}
		switch char {
		case '0':
			value = 0
		case '1':
			value = 1
		case '2':
			value = 2
		case '3':
			value = 3
		case '4':
			value = 4
		case '5':
			value = 5
		case '6':
			value = 6
		case '7':
			value = 7
		case '8':
			value = 8
		case '9':
			value = 9
		case 'A':
			value = 10
		case 'B':
			value = 11
		case 'C':
			value = 12
		case 'D':
			value = 13
		case 'E':
			value = 14
		case 'F':
			value = 15
		default:
			continue
		}

		result = result*16 + value
	}

	return result
}

func main() {
	hexValue := "1E"
	byteValue := HexToByte(hexValue)
	fmt.Println("Décimal (Byte):", byteValue)
}

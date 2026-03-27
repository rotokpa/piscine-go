package piscinego

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var converter string
	fmt.Print("Enter converter: (bin, hex, dec): " )
	fmt.Scan(&converter)

	var convertnum string 
	fmt.Print("enter number to be converted: ")
	fmt.Scan(&convertnum)
	
	switch convertnum {
	case "dec":
		hex := strings.ToUpper(strconv.FormatInt(dec, 16, 64))
		fmt.Println("Hex: ", hex)
		
		binary := strconv.FormatInt(dec, 2, 64)
		fmt.Println("Binary: ", binary)
	default:
		fmt.Print("is not a valid decimal")
	}
	
	func convertBin(bin string) (int64 error) {
		val, err := strconv.ParseInt(bin, 2, 64)
		if err != nil {
			fmt.Print("is not valid binary: ", bin)
		}
		return 
	}

	func converthex(hex string) (int64 error) {
		val, err := strconv.ParseInt(hex, 16, 64)
		return val, err
	}
	
}
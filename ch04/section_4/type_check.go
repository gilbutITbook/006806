package main

import "fmt"

func checkType(v interface{}) {
	switch v.(type) {
	case bool:
		fmt.Printf("%t is a bool\n", v)
	case int, int8, int16, int32, int64:
		fmt.Printf("%d is an int\n", v)
	case uint, uint8, uint16, uint32, uint64:
		fmt.Printf("%d is an unsigned int\n", v)
	case float64:
		fmt.Printf("%f is a float64\n", v)
	case complex64, complex128:
		fmt.Printf("%f is a complex\n", v)
	case string:
		fmt.Printf("%s is a string\n", v)
	case nil:
		fmt.Printf("%v is nil\n", v)
	default:
		fmt.Printf("%v is unknown type\n", v)
	}
}

func main() {
	checkType(3)             // 3 is an int
	checkType(1.5)           // 1.500000 is a float64
	checkType(complex(1, 5)) // (1.000000+5.000000i) is a complex
	checkType(true)          // true is a bool
	checkType("s")           // s is a string
	checkType(struct{}{})    // {} is unknown type
	checkType(nil)           // <nil> is nil
}

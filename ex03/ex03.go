package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// bitOperations()
	floatOper()
}

func bitOperations() {
	var x, y, z uint8

	x = 9
	y = 28
	z = x
	fmt.Println("Bit Operations")
	fmt.Printf("%.8b\n%8b\n%8b\n\n", x, x, ^x)
	fmt.Printf("%.8b\n%8b\n%8b\n\n", x, y, x&y)
	fmt.Printf("%.8b\n%8b\n%8b\n\n", x, y, x|z)
	fmt.Printf("%.8b\n%8b\n%8b\n\n", x, y, x&^x)
}

func floatOper() {
	var defaultFloat float32
	floatVar := 5.5

	fmt.Println("defaultFloat   = ", defaultFloat, "\n")
	fmt.Printf("floatVar (%s)  = %v\n", reflect.TypeOf(floatVar), floatVar)
	fmt.Println("\t-----------------------------------------------------------\n\n")

	fmt.Println("MAX float32    = ", math.MaxFloat32)                   // 3.4028234663852886e+38
	fmt.Println("MIN float32    = ", math.SmallestNonzeroFloat32, "\n") // 1.401298464324817e-45
	fmt.Println("MAX float64    = ", math.MaxFloat64)                   // 1.7976931348623157e+308
	fmt.Println("MIX float64    = ", math.SmallestNonzeroFloat64, "\n") // 5e-324
}

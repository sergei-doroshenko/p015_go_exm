package main

import (
	"fmt"
)

const SingelLineString = "Single line constant"

const (
	MultiLineFirst  = 1       // 1
	MultiLineSecond           // 1
	MultiLineThird  = "third" // "third"
	MultiLineFourth           // "third"
)

const (
	IotaFirst  = iota //  0
	IotaSecond        // 1
	IotaThird         // 2
	IotaFourth        // 3
)

const (
	IotaUseCaseFirst  = 1 << iota //  1
	IotaUseCaseSecond             // 2
	IotaUseCaseThird              // 4
	IotaUseCaseFourth             // 8
)

const (
	IotaUseCase2First  = iota*3 - 2 //  -2
	IotaUseCase2Second              // 1
	IotaUseCase2Third               // 4
	IotaUseCase2Fourth              // 7
)

func main() {
	fmt.Println("Single line constant: ", SingelLineString, "\n")
	fmt.Println("MultiLineFirst: ", MultiLineFirst)
	fmt.Println("MultiLineSecond: ", MultiLineSecond)
	fmt.Println("MultiLineThird: ", MultiLineThird)
	fmt.Println("MultiLineFourth: ", MultiLineFourth, "\n")

	fmt.Println("IotaFirst: ", IotaFirst)
	fmt.Println("IotaSecond: ", IotaSecond)
	fmt.Println("IotaThird: ", IotaThird)
	fmt.Println("IotaFourth: ", IotaFourth, "\n")

	fmt.Println("IotaUseCaseFirst: ", IotaUseCaseFirst)
	fmt.Println("IotaUseCaseSecond: ", IotaUseCaseSecond)
	fmt.Println("IotaUseCaseThird: ", IotaUseCaseThird)
	fmt.Println("IotaUseCaseFourth: ", IotaUseCaseFourth, "\n")

	fmt.Println("IotaUseCase2First: ", IotaUseCase2First)
	fmt.Println("IotaUseCase2Second: ", IotaUseCase2Second)
	fmt.Println("IotaUseCase2Third: ", IotaUseCase2Third)
	fmt.Println("IotaUseCase2Fourth: ", IotaUseCase2Fourth, "\n")
}

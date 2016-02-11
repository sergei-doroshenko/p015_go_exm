package main

import (
	"fmt"
	// whole path from src folder -> go install
	"github.com/sergei-doroshenko/p015_go_exm/ex01/util"
)

func main() {
	fmt.Println("Hello, Sergei!!!")
	fmt.Println("System type: " + util.GetOsType())
}

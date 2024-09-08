package main

import (
	"fmt"
	"go_prg/package1"
)

func main() {
	// fmt.Println("Hello, World!")
	package1.PrintHelloWorld()
	fmt.Print(new_normal_struct())
	anonymous_struct := struct {
		k  int
		mn float64
		nn string
	}{
		k:  1,
		mn: 1.0,
		nn: "1",
	}
	fmt.Print(anonymous_struct)
}

type normal_struct struct {
	k  int
	f  float32
	mn float64
	nn string
	nf bool
}

func new_normal_struct() normal_struct {
	nor_instance := normal_struct{}
	nor_instance.k = 1
	nor_instance.f = 1.0
	nor_instance.mn = 1.0
	nor_instance.nn = "1"
	nor_instance.nf = false
	if nor_instance.nf {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
	// fmt.Println(nor_instance)
	return nor_instance
}

// defining anonymous struct

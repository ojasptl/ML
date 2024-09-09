package main

import (
	"errors"
	"fmt"
	"go_prg/interface1"
	"go_prg/package1"
)

func main() {
	// fmt.Println("Hello, World!")
	a := []string{"a", "b", "c"}
	for _, name := range a {
		// for loop for list gives index and value both
		fmt.Println(name)
	}
	fmt.Print(interface1.Aggregate(1, 2, 3, interface1.Add))
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
	ghn, _, _ := interface1.Interface3()
	fmt.Print(ghn)
	maker := engine{}
	maker.brand = "Toyota"
	maker.model = "Corolla"
	maker.power = 1000
	fmt.Println(maker)
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
	nested_struct()
	// used strfuct with new method

	engine := engine{power: 1000}
	fmt.Print(engine.getPower())
	return nor_instance
}

// defining anonymous struct

// embedded struct

type car struct {
	brand string
	model string
}

type engine struct {
	power int
	car
}

func (e engine) getPower() int {
	return e.power
}

func nested_struct() {
	// nested struct
	crt := car{
		brand: "Toyota",
		model: "Corolla",
	}
	fmt.Println(crt)
}
func main1(a int, b string) (string, error) {
	h := make(map[string]map[string]int)
	m, ok := h["a"]["b"]
	if m == a {
		fmt.Print(m, ok)

		return " " + b, nil
	}
	// abn := 1
	fmt.Print(m, ok)
	return " ", errors.New("Error")
}

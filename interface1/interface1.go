package interface1

import (
	// "errors"
	"errors"
	"fmt"
)

func getUser(in bool) (string, error) {
	if in {
		return "", errors.New("Error")
	}
	return "User", nil
}
func Interface3() (string, []int, error) {
	user, err := getUser(true)
	if err != nil {
		fmt.Println(err)
		return "0", new(), err
	}
	// return 1, nil
	fmt.Println(user)
	return " ", new(), err
}

func new() []int {
	primes := []int{2, 3, 5, 7, 11, 13}
	// abc []byte
	// fmt.Println(abc)
	primes = append(primes, 17)
	x := make_map()
	fmt.Print(x)
	y := make_map_1()
	fmt.Print(len(y))
	primes[1] += 1
	return primes
}

func make_map() map[string]int {
	// var m map[string]int
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	return m
}

func make_map_1() map[string]int {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	delete(m, "two")
	elem, ok := m["two"]
	fmt.Print(elem, ok)
	mgm := user{
		name: "John",
		age:  30,
	}
	elf := make(map[string]user)
	elf["one"] = mgm
	str, err := new123(elf)
	fmt.Print(str, err)

	// if elem exists in two then ok is true else false
	return m
}

// s ...int -> means a slice of array meaning same as
// s []int
// func sum(s ...int) int {
// 	n := 0
// 	for _, v := range s {
// 		n += v
// 	}
// 	return n
// }
// func main() {
// 	fmt.Print(sum(1, 2, 3, 4, 5))
// 	a []int{1, 2, 3, 4, 5}
// 	fmt.Print(sum(a...))
// }

type user struct {
	name string
	age  int
}

func new123(user1 map[string]user) (deleted bool, err error) {
	if _, ok := user1["one"]; !ok {
		delete(user1, "one")
		return true, nil
	}
	return false, errors.New("Error")
}

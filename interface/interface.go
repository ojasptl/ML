package interface1

import (
	// "errors"
	"fmt"
)

func getUser() (string, error) {
	return "1", nil
}
func interface1() {
	user, err := getUser()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user)
}

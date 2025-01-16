package main

import (
	"fmt"

	"github.com/nishi7409/goblox/lib/users"
)

func init() {

}

func main() {
	const username = "-"

	resp, err := users.GetIDFromUsername(username)
	fmt.Printf("Username: %v => %v\nError => %v\n", username, resp, err)
}

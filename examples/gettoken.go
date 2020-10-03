package main

import (
	"fmt"
	"os"

	"github.com/ProZsolt/elmuemasz"
)

func main() {
	username := os.Getenv("ELMU_USERNAME")
	password := os.Getenv("ELMU_PASSWORD")
	user := elmuemasz.User{
		Username: username,
		Password: password,
	}

	ts, err := elmuemasz.NewTokenSource(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	t, err := ts.Token()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t.AccessToken)
}

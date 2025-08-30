package main

import (
	"TestProject/cmd"
	"fmt"
)

func main() {
	if err := cmd.ExecuteProgramm(); err != nil {
		fmt.Println(err)
	}
}

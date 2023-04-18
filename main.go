package main

import (
	"fmt"
	"mvc/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	fmt.Println("Hello")
}

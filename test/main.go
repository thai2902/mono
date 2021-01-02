package main

import (
	"fmt"

	"github.com/thai2902/mono/auth"
)

func main() {
	hello := auth.GetToken()
	fmt.Print(hello)
}

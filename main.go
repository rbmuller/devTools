package main

import (
	"fmt"

	devErrors "github.com/rbmuller/devtools/errors"
)

func main() {
	x := devErrors.NewError("Test new error")

	fmt.Println(x)

}

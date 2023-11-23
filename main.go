package main

import (
	"fmt"

	devErrors "github.com/devtools/errors"
)

func main() {
	x := devErrors.NewError("Test new error")

	fmt.Println(x)

}

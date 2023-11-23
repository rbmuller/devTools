package main

import (
	"errors"
	"fmt"

	devErrors "github.com/rbmuller/devtools/errors"
	devFilters "github.com/rbmuller/devtools/filters"
	devProcess "github.com/rbmuller/devtools/process"
)

func main() {
	//Showing how to better debug with devErrors
	x := devErrors.NewError(errors.New("print here the error message from your function"))
	fmt.Println(x)

	//Showing how to deduplicate an array
	TestArray := []string{"A", "A", "B", "C"}

	y := devFilters.DeduplicateArray(TestArray)
	fmt.Println(y)

	//Showing Epoch to Timestamp
	z, ok := devProcess.EpochToTimestamp(1684624830053)
	if !ok {
		fmt.Println(devErrors.NewError(errors.New("invalid Epoch format")))
	}
	fmt.Println(z)

}

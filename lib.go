package B

import (
	"fmt"

	"github.com/jun06t/C"
)

const version = "1.2.0"

func Func() string {
	fmt.Println(C.Func())
	return version
}

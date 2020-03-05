package mylib
import (
	"fmt"
	"runtime"
)

// DisplayVersion ...
// DisplayVersion ...
func DisplayVersion(){
	fmt.Println(runtime.Version())
}
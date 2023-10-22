// toma una desicion en base a una valor
package main

import (
	"fmt"
	"runtime"
)

func main() {
	os := runtime.GOOS
	switch os {
	case "linux":
		fmt.Println("Go run ->Linux")
	case "windows":
		fmt.Println("go run ->Windows")
	case "darwin":
		fmt.Println("go run ->macOS")
	default:
		fmt.Println("go run ->Otros OS")

	}
}

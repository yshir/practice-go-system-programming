package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintf(os.Stdout, "int = %d, str = %s, float = %f", 1, "a", 1.0)
}

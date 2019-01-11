package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprint(os.Stdout, "a", "b", "c")
	fmt.Fprintln(os.Stdout, "a", "b", "c")
	fmt.Fprintf(os.Stdout, "%s|%s|%s", "a", "b", "c")
}

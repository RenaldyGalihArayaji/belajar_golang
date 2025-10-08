package lines

import "fmt"

// Line prints a separator and the provided text.
// Exported so it can be used from other packages.
func Line(text string) {
	fmt.Println("====================================")
	fmt.Println(text)
	fmt.Println("====================================")
}

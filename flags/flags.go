package flags

import (
	"flag"
	"fmt"
)

func Flags() {
	num := flag.Int("n", 0, "# of iteration")
	entry := flag.String("t", "empty", "New task")
	flag.Parse()

	n := *num
	//prints the type
	fmt.Printf("var type: %T\n", n)

	for i := 0; i < n; i++ {
		// 1:=0
		// for i < n {i++}
		fmt.Printf("hello")
	}
	if *entry != "empty" {
		fmt.Printf("%v\n", *entry)
	}
}

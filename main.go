package main

import (
	"flag"
	"fmt"
)

//Go’s structs are typed collections of fields. They’re useful for grouping data together to form records.
//Go does not have optional type (remember that go a declared var must be used)
type Entry struct {
	msg  string
	date int
}

func main() {
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

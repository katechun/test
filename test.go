package main

import (
	"fmt"
	"github.com/katechun/test/a"
	"github.com/katechun/test/b"
)

func main() {
	m := a.Ss{
		Id:1,
		Name:"katechun",
	}
	b1:=b.Ss(m)


	fmt.Println(b1)
}

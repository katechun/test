package main

import (
	"fmt"
	"os"

)

func main() {
	f:=createFile("/tmp/defe.txt")
	defer closeFile(f)
	writeFile(f)
}


func createFile(p string) *os.File {
	fmt.Println("creating")
	f,err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f

}

func closeFile(f *os.File){
	fmt.Println("closing")
	f.Close()
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f,"data")
}
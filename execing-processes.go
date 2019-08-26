package main

import (
	"os"
	"os/exec"
	"syscall"
)

func check(e error) {
	if e!= nil {
		panic(e)
	}
}

func main() {
	binary,lookErr := exec.LookPath("ls")
	check(lookErr)

	args := []string{"ls","-a","-l","-h"}

	env := os.Environ()

	execErr := syscall.Exec(binary,args,env)
	check(execErr)
}

package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		execute()
	}
}

func execute() {
	out, err := exec.Command("ls").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println("Command [ls] Successfully Executed")
	output := string(out[:])
	fmt.Println(output)

	out, err = exec.Command("pwd").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println("Command [pwd] Successfully Executed")
	output = string(out[:])
	fmt.Println(output)

	// Passing in Arguments
	out, err = exec.Command("ls", "-ltr").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println("Command [ls -ltr] Successfully Eexecuted")
	output = string(out[:])
	fmt.Println(output)
}

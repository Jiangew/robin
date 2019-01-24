package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("------ Reading and Writing Files ------")

	mydata := []byte("All the data i wish to write to a file\n")

	// WriteFile method returns an error if unsuccessful
	err := ioutil.WriteFile("myfile.data", mydata, 0777)
	// handle error
	if err != nil {
		fmt.Println(err)
	}

	data, err := ioutil.ReadFile("myfile.data")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(data))

	f, err := os.OpenFile("myfile.data", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString("New data that wasn't there originally\n"); err != nil {
		panic(err)
	}

	data, err = ioutil.ReadFile("myfile.data")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(data))
	fmt.Println("------ Reading and Writing Files End ------")
}

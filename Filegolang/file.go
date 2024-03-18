package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	content := "THE FILE IS BEEN CREATED HERE !!!"
	fmt.Println("File on load")

	file, err := os.Create("./createdFile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	checkNilError(err)

	fmt.Println(length)

	defer file.Close()

	readFile("./createdfile.txt")
}

func readFile(filename string) {

	dbyte, err := ioutil.ReadFile(filename)
	checkNilError(err)
	checkNilError(err)

	fmt.Println("THE FILE DATA Byte IS \n\n", dbyte)

	fmt.Println("FILE CONTENT : \n", string(dbyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

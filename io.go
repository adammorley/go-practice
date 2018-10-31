package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func simpleReadFile() {
	b, e := ioutil.ReadFile("file")
	if e != nil {
		panic(e)
	}
	fmt.Println(b)
}

func chunkedReadFile() {
	f, e := os.Open("file")
	if e != nil {
		panic(e)
	}
	defer f.Close()

	var b []byte = make([]byte, 4)
	for e == nil {
		_, e = f.Read(b)
		fmt.Println(b)
	}
    if e != io.EOF {
        panic(e)
    }
}

func chunkedWriteFile() {
    f, e := os.OpenFile("file", os.O_WRONLY | os.O_TRUNC | os.O_CREATE, 0644)
    if e != nil {
        panic(e)
    }
    defer f.Close()

    var b []byte = make([]byte, 0, 16)
    var s string = "this is a string"
    b = []byte(s)
    n, e := f.Write(b)
    if e != nil {
        panic(e)
    } else if n != 16 {
        panic(n)
    }
    e = f.Sync()
    if e != nil {
        panic(e)
    }
}

func main() {
	simpleReadFile()
	chunkedReadFile()
    chunkedWriteFile()
}

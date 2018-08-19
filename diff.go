package main

import "bytes"
import "io"
import "io/ioutil"
import "os"
import "fmt"

var file0 string = "file0"
var file1 string = "file1"

func firstPass() {
    var fileData0 []byte = make([]byte, 0, 100) // optimization: make byte buffer size of file using io library
    var e error
    fileData0, e = ioutil.ReadFile(file0) // has to read whole file into memory!
    if e != nil {
        panic(e)
    }
    var fileData1 []byte = make([]byte, 0, 100)
    fileData1, e = ioutil.ReadFile(file1)
    if e != nil {
        panic(e)
    }
    // comparison is going to work because the bytes will be the same (assuming endianness here)
    // first a simple comparison: the length!
    if len(fileData0) != len(fileData1) {
        fmt.Println("files are different lengths")
        os.Exit(1)
    }
    for i := 0; i < len(fileData0); i++ {
        if fileData0[i] != fileData1[i] {
            fmt.Println("files differ at location", i)
            fmt.Println("f0: ", string(fileData0[i]), "f1: ", string(fileData1[i]))
        }
    }
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func secondPass() {
    var f0, f1 *os.File
    var e error
    f0, e = os.Open(file0)
    check(e)
    f1, e = os.Open(file1)
    var f0i, f1i os.FileInfo
    f0i, e = f0.Stat()
    check(e)
    f1i, e = f1.Stat()
    check(e)
    if f0i.Size() != f1i.Size() {
        fmt.Println("files are different sizes")
    }
    var f0d, f1d []byte = make([]byte, 64), make([]byte, 64)
    var d bool = false;
    for {
        var e0, e1 error
        _, e0 = f0.Read(f0d)
        _, e1 = f1.Read(f1d)
        if e0 != e1 {
            fmt.Println(e0, e1)
            panic("somethings wrong, should not get here")
        }
        if bytes.Compare(f0d, f1d) != 0 {
            fmt.Println("files differ")
            fmt.Println(string(f0d), string(f1d))
            d=true
        }
        if e0 == io.EOF && d {
            os.Exit(1)
        } else {
            os.Exit(0)
        }
    }
}

func main() {
    //firstPass()
    secondPass()
}

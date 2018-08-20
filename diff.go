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
	var d bool = false
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
			d = true
		}
		if e0 == io.EOF && d {
			os.Exit(1)
		} else {
			os.Exit(0)
		}
	}
}

func openFile(fileName string) (fd *os.File) {
	var e error
	fd, e = os.Open(fileName)
	check(e)
	return
}

func statFile(file *os.File) (fi os.FileInfo) {
	var e error
	fi, e = file.Stat()
	check(e)
	return
}

// returns true if files are the same size
func notSizeSame(file0 *os.File, file1 *os.File) bool {
	if statFile(file0).Size() != statFile(file1).Size() {
		return true
	}
	return false
}

func bytesDiffer(d0 []byte, d1 []byte) bool {
	if bytes.Compare(d0, d1) == 0 {
		return false
	}
	return true
}

const FBufSz uint = 4096 // add dynamic page size discovery and maybe some readahead buffering

func thirdPass() {
	f0 := openFile(file0)
	f1 := openFile(file1)
	if notSizeSame(f0, f1) {
		fmt.Println("file sizes are different")
		os.Exit(0)
	}
	var d0, d1 []byte = make([]byte, FBufSz), make([]byte, FBufSz)
	for {
		var e0, e1 error
		_, e0 = f0.Read(d0)
		_, e1 = f1.Read(d1)
		if e0 != e1 {
			// the files are the same size, so should not get here
			panic("something is wrong!")
		} else if e0 == io.EOF {
			os.Exit(0)
		} else if bytesDiffer(d0, d1) {
			fmt.Println("files differ, f0: ", string(d0), " f1: ", string(d1))
		}
	}
}

func main() {
	//firstPass()
	//secondPass()
	thirdPass()
}

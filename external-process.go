package main

import "bytes"
import "fmt"
import "os"
import "os/exec"

func main() {
    var c *exec.Cmd
    c = exec.Command("/bin/false")
    var o bytes.Buffer
    c.Stdout = &o
    var e error
    e = c.Run()
    if _, ok := e.(*os.PathError); ok {
        fmt.Println("could not find command")
    } else if _, ok := e.(*exec.ExitError); ok {
        fmt.Println("error executing program")
    } else if e == nil {
        fmt.Println("program ran and exited cleanly")
    }
    fmt.Println(e)
}

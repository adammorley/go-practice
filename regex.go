package main

import(
"fmt"
"regexp"
)

func main() {

var text string = "the cow jumped over the moon"

var r *regexp.Regexp=regexp.MustCompile(`(\w+)\s+jumped`)
fmt.Println(r.MatchString(text))
fmt.Println(r.FindStringSubmatch(text)[1])
}

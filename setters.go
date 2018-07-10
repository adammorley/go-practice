package main

import "errors"
import "fmt"

type thing struct {
	owner, name string
}

func (t *thing) Owner(x ...string) (string, error) {
	fmt.Printf("%T\n", x)
	fmt.Printf("%d\n", len(x))
	if len(x) == 0 {
		return t.owner, nil
	} else if len(x) == 1 {
		t.owner = x[0]
		return t.owner, nil
	} else {
		return "", errors.New("too many args")
	}
}

func main() {
	var t thing
	o, _ := t.Owner()
	fmt.Println(o)
	o, _ = t.Owner("alice")
	fmt.Println(o)
	o, _ = t.Owner()
	fmt.Println(o)
	_, e := t.Owner("bob", "susie")
	fmt.Println(e)
}

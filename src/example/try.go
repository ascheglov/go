package main

import "fmt"
import "errors"

func h() (int, error) {
	return 0, errors.New("OMG")
}

func g() (__err error) {
	x := __try h() // assignment
	fmt.Println("not reached", x)
	return
}

func f() (__err error) {
	__try g() // statement
	fmt.Println("not reached")
	return
}

func main() {
	fmt.Println("start")
	err := f()
	fmt.Println("done, error: ", err)
}

// Expected output:
//
// $ go run example/try.go
// start
// done, error:  OMG
//

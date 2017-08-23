package main

import "fmt"
import "errors"

func g() error {
	return errors.New("OMG")
}

func f() error {
	fmt.Println("before try")
	__try g() // MAGIC!
	fmt.Println("not reached")
	return nil
}

func main() {
	err := f()
	fmt.Println("done, error: ", err)
}

// Expected output:
//
// $ go run example/try.go
// before try
// done, error:  OMG
//

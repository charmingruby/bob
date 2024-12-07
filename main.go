package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	e := exec.Command("make", "new-mig", "NAME=foo")
	var out bytes.Buffer
	e.Stdout = &out
	err := e.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Output: %q\n", out.String())
}

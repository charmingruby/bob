package main

import (
	"fmt"

	"github.com/charmingruby/gentoo/config"
)

func main() {
	config, err := config.New("./dummy")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", config)
}

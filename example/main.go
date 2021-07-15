package main

import (
	"fmt"

	"github.com/isakgranqvist2021/env_vars/env_vars"
)

func main() {
	vars := env_vars.Get()

	for _, v := range vars {
		fmt.Printf("%s = %s\n", v.Key, v.Value)
	}
}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Variable struct {
	Key   string
	Value string
}

func main() {
	var variables []Variable
	bytes, err := ioutil.ReadFile(".env")

	if err != nil {
		log.Fatal(err)
	}

	reader := strings.NewReader(string(bytes))
	n := len(string(bytes))

	var t Variable

	for i := 0; i < n; i++ {
		var s string
		fmt.Fscanln(reader, &s)
		if len(s) > 0 {
			if i%2 == 0 {
				t.Key = s
			} else {
				t.Value = s
				variables = append(variables, t)
			}
		}
	}

	for _, v := range variables {
		fmt.Printf("%s : %s\n", v.Key, v.Value)
	}
}

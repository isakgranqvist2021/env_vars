package env_vars

import (
	"io/ioutil"
	"log"
	"strings"
)

type variable struct {
	Key   string
	Value string
}

func convStr(v string, c chan variable) {
	vals := strings.Split(v, "=")

	c <- variable{
		Key:   strings.Trim(strings.Trim(vals[0], " "), "\n"),
		Value: strings.Trim(strings.Trim(vals[1], " "), "\n"),
	}
}

func Get() []variable {
	var variables []variable

	bytes, err := ioutil.ReadFile(".env")

	if err != nil {
		log.Fatal(err)
	}

	c := make(chan variable)

	for _, v := range strings.SplitAfter(string(bytes), "\n") {
		go convStr(v, c)
		variables = append(variables, <-c)
	}

	return variables
}

package main

import (
	flag "github.com/ogier/pflag"
)

// flags
var (
	user string
)

func init() {
	flag.StringVarP(&user, "user", "u", "Search Users")
}

func main() {
	flag.Parse()
}
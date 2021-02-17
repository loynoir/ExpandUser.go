package main

import (
	"fmt"
	"github.com/loynoir/ExpandUser.go"
)

func main() {
	x, _ := ExpandUser.ExpandUser("~/~/foo/bar/")
	fmt.Println(x)

	TestCases := []string{
		"~",
		"~/",
		"~/~",
		"~/~/",
		"~/foo",
		"~/foo/bar",
		"~/foo/bar/",
		"~foo",
		"foo",
		"foo~",
		"foo~bar",
	}
	for _, el := range TestCases {
		_path, _ := ExpandUser.ExpandUser(el)
		fmt.Println(el + " => " + _path)
	}
}

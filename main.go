package main

import (
	"fmt"

	"github.com/didzl/bank_dic/mydict"
)


func main() {
	dictionary:=mydict.Dictionary{"first": "first word"}
	word := "hello"
	definition :="greeting"

	err :=dictionary.Add(word, definition)
	if err !=nil {
		fmt.Println(err)
	}
	hello, _ := dictionary.Search(word)
	fmt.Println(hello)
	err2:= dictionary.Add(word, definition)
	if err2 !=nil {
		fmt.Println(err2)
	}

}
package main

import (
	"fmt"

	"github.com/didzl/bank_dic/mydict"
)


func main() {
	dictionary:=mydict.Dictionary{"first": "first word"}
	baseWord := "hello"
	definition :="greeting"

	dictionary.Add(baseWord, definition)
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word,err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(word)
	}
}
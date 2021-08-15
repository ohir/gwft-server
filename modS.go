package helloS

import "fmt"

import (
	proto "example.com/mP"
)

func Id() string { return "Server" }
func Hello() { fmt.Printf("%s <= %s\n", Id(), proto.Id()) }
func Say(m, s string) { fmt.Printf("%s <= %s\n",m,s) }

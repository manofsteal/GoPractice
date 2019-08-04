package main

// built-in package have short name
import (
	"fmt"
	"time"
)

// this will find package in import path (files name does not matter)
// then assign channel =  found_package
// every files in same directory have to contain same package
// in convention, directory (import path) is same as package name
import "practice/routine/channel" // recommmend
// import "./channel" // worked but not recommmend

type aType struct {
	data int
}

func (t *aType) change() {
	t.data = 10
}

func main() {

	ipc := channel.Make()

	s2ns := func(s int) time.Duration {
		return time.Duration(s * 2000000000)
	}

	go func() {

		signal := uint8(0)
		for true {
			time.Sleep(s2ns(1))
			ipc.Push(channel.Message{signal, signal + 1})

			signal += 1
		}

	}()

	go func() {
		for true {
			wait := ipc.Pull()

			fmt.Println("message", len(wait), wait)
		}

	}()

	t := aType{
		data: 11,
	}

	t.change()

	fmt.Println(t.data)

	for {
		time.Sleep(s2ns(10))
	}

}

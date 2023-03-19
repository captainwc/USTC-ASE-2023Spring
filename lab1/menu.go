package main

import "fmt"

func main() {
	for {
		var cmd string
		fmt.Printf("==> Please give your command $ ")
		fmt.Scanln(&cmd)
		if cmd == "help" {
			fmt.Println(" ->hello: give you a greeting from the author")
			fmt.Println(" ->exit: close this menu")
			fmt.Println(" ->help: show this help text")
		} else if cmd == "hello" {
			fmt.Println("How is everything going?")
		} else if cmd == "exit" {
			fmt.Println("exit")
			break
		} else {
			fmt.Println("ERROR: Illeagl Command!\n\tPlease cheak out your input, or get some help by type \"help\"")
		}
	}
}

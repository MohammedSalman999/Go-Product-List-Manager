package main

import "fmt"

func main() {
	store := list{
		{title: "helo", price: 150, released: isTimeStamp("")},
		{title: "mobidick", price: 150, released: isTimeStamp(0)},
	}
	store.discount(0.5)
	fmt.Print(store)

}
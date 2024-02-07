package main

import "fmt"

type product struct {
	title    string
	price    money
	released timestamp
}

func (p *product) String() string {
	s:= fmt.Sprintf("%s : %s : %s", p.title, p.price ,p.released)
	return s
}

func (p *product) discount(ratio float64) {
	p.price *= money(1- ratio)
}
package main

type book struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Quatity int    `json:"quality"`
}

type client struct {
	ID   int
	name string
	cpf  string
}

type loan struct {
	Books    []book `json:"books"`
	LoanDate string `json:"loandate"`
	Client   client `json:"client"`
}

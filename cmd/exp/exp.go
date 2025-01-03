package main

import (
	"html/template"
	"os"
)

func main() {
	t, err := template.New("teste").Parse("<h1>Título</h1>")
	if err != nil {
		panic(err)
	}
	t.Execute(os.Stdout, nil)
}

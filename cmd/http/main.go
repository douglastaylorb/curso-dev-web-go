package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func noteList(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/templates/home.html")

	if err != nil {
		http.Error(w, "Houve um erro na renderização da página", http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)
}

func noteView(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID é obrigatório", http.StatusBadRequest)
		return
	}
	note := `
		<div>
			<h3> Nota %s </h3>
			<p> Detalhes da nota %s </p>
		</div>
	`

	fmt.Fprint(w, note, id)
}

func noteCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprint(w, "Criando uma nova anotação")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", noteList)
	mux.HandleFunc("/note/view", noteView)
	mux.HandleFunc("/note/create", noteCreate)

	fmt.Println("Servidor rodando na porta: 5000")
	http.ListenAndServe(":5000", mux)

}

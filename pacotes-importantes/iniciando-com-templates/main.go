package main

import (
	"html/template"
	"net/http"
	"strings"
)

// Iniciando com templates: Aulas basicas em um arquivo só.
// Templates com arquivo externo
// Templates com Web Server
// Compondo templates
// Mapeando funções nos templates

type Anime struct {
	Nome string
	Ano  int
	Nota float64
}

type Animes []Anime

func main() {

	listaAnimes := listComAnimes()

	templates := []string{
		"header.html",
		"anime.html",
		"footer.html",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.New("anime.html")
		t.Funcs(template.FuncMap{"ToUpper": ToUpper})
		t, err := t.ParseFiles(templates...)
		if err != nil {
			panic(err)
		}

		err = t.Execute(w, listaAnimes)
		if err != nil {
			panic(err)
		}

	})

	http.ListenAndServe(":8080", nil)

}

func listComAnimes() Animes {
	anime01 := Anime{
		Nome: "Evangelion",
		Ano:  1995,
		Nota: 10.0,
	}

	anime02 := Anime{
		Nome: "Bleach",
		Ano:  2004,
		Nota: 8.0,
	}

	anime03 := Anime{
		Nome: "Shigatsu wa Kimi no Uso",
		Ano:  2014,
		Nota: 10.0,
	}

	anime04 := Anime{
		Nome: "Kaguya-sama wa Kokurasetai: Ultra Romantic",
		Ano:  2022,
		Nota: 10.0,
	}

	anime05 := Anime{
		Nome: "Fate/stay night: Unlimited Blade Works",
		Ano:  2014,
		Nota: 9.0,
	}
	return Animes{anime01, anime02, anime03, anime04, anime05}
}

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

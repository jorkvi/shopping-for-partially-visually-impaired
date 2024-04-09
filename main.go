package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type PageData struct {
	Title        string
	ImageURL     string
	Avatar       string
	Parduotuve   string
	Atsiliepimai string
	Nustatymai   string
	Ataskaita    string
	Apie         string
	Lupa         string
	Rodykle      string
	Ikona        string
	Bruksniai    string
	Diagramos    string
	Prekes       string
	Minus        string
	Plius        string
	Keisti       string
}

func main() {
	http.HandleFunc("/homePage", homeHandler)
	http.HandleFunc("/login", HomePage)
	http.HandleFunc("/logine", Login)
	//http.HandleFunc("/dashboard", Dashboard)
	http.HandleFunc("/logo.png", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "logo.png")
	})

	http.Handle("/gallery/", http.StripPrefix("/gallery/", http.FileServer(http.Dir("gallery"))))
	http.HandleFunc("/pagrindinis", h1)

	http.HandleFunc("/delete", deleteHandler)

	http.HandleFunc("/edit", editHandler)
	http.HandleFunc("/edit1", updateHandler)
	http.HandleFunc("/insert", h2)
	http.HandleFunc("/insert1", insertHandler)
	http.HandleFunc("/search", searchHandler)

	fmt.Println("Serveris veikia adresu: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
	fmt.Println("hello world")
	//	http.Handle("/gallery/", http.StripPrefix("/gallery", http.FileServer(http.Dir("./gallery"))))

	log.Fatal(http.ListenAndServe(":8000", nil))

}

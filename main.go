package main

import (
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
	UserData     *User
}

func main() {
	http.HandleFunc("/homePage", homeHandler)
	http.HandleFunc("/login", HomePage)
	http.HandleFunc("/logine", Login)
	http.HandleFunc("/logo.png", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "logo.png")
	})

	http.Handle("/gallery/", http.StripPrefix("/gallery/", http.FileServer(http.Dir("gallery"))))

	// Nustatome middleware, kuris tikrins, ar vartotojas yra prisijungęs prie kiekvieno užklausos

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, "session-name")
		if session.Values["LoggedIn"] == nil {
			// Jei vartotojas nėra prisijungęs, jį peradresuojame į prisijungimo puslapį
			http.Redirect(w, r, "/homePage", http.StatusSeeOther)
			return
		}

		// Jei vartotojas yra prisijungęs, leidžiame prieinamumą pagal reikalavimus
		switch r.URL.Path {
		case "/pagrindinis":
			h1(w, r)
		case "/delete":
			deleteHandler(w, r)
		case "/edit":
			editHandler(w, r)
		case "/edit1":
			updateHandler(w, r)
		case "/insert":
			h2(w, r)
		case "/insert1":
			insertHandler(w, r)
		case "/search":
			searchHandler(w, r)
		case "/admin":
			handlerFunc(w, r)
		case "/adminEdit":
			admindata(w, r)
		case "/tikrinti":
			barCodeHandler(w, r)
		default:
			// Jei vartotojas nukreiptas į kitą URL, tai peradresuojame į pagrindinį puslapį
			http.Redirect(w, r, "/homePage", http.StatusSeeOther)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

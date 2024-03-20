package main

import (
	"database/sql"
	"fmt"
	"html/template"
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

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("hello world")
	//	http.Handle("/gallery/", http.StripPrefix("/gallery", http.FileServer(http.Dir("./gallery"))))

	http.HandleFunc("/", h1)

	http.Handle("/gallery/", http.StripPrefix("/gallery/", http.FileServer(http.Dir("gallery"))))

	http.HandleFunc("/delete", deleteHandler)

	http.HandleFunc("/edit", editHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))

}

func editHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully connected to the database!")

	id := r.FormValue("id")

	//query := fmt.Sprintf("SELECT * FROM `produktai` WHERE produktai.id = '%s'", id)

	//	if _, err := db.Exec(query); err != nil {
	//		http.Error(w, err.Error(), http.StatusInternalServerError)
	//		return
	//	}

	type Produktas struct {
		Pavadinimas string
		Kodas       int
		Parduotuve  string
		Plius       string
		Minus       string
		Keisti      string
		ID          int
	}
	var records []Produktas
	query := fmt.Sprintf("SELECT * FROM `produktai` WHERE produktai.id = '%s'", id)

	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var record Produktas
		record.Minus = "gallery/circle-minus-1.png"
		record.Plius = "gallery/Circled_plus.png"
		record.Keisti = "gallery/keisti.png"
		if err := rows.Scan(&record.Kodas, &record.Pavadinimas, &record.Parduotuve, &record.ID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		records = append(records, record)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//}
	data := PageData{
		Title:        "Example Page",
		ImageURL:     "gallery/25694.png", // URL to your image
		Avatar:       "gallery/3837171.png",
		Parduotuve:   "gallery/1413908.png",
		Atsiliepimai: "gallery/787610-200.png",
		Nustatymai:   "gallery/setting.png",
		Ataskaita:    "gallery/1268.png",
		Apie:         "gallery/question.png",
		Lupa:         "gallery/lupa.png",
		Rodykle:      "gallery/rodykle1.png",
		Ikona:        "gallery/ikona.png",
		Bruksniai:    "gallery/bruksniai.png",
		Diagramos:    "gallery/diagramos.bmp",
		Prekes:       "gallery/prekes.png",
		Plius:        "gallery/Circled_plus.png",
		Minus:        "gallery/circle-minus-1.png",
		Keisti:       "gallery/keisti.png",
	}
	tmpl := template.Must(template.ParseFiles("redagavimas.html"))
	//fmt.Fprintf(w, "<form method='post'><input type='text' name='data'><input type='submit' value='Siųsti'></form>")
	if err := tmpl.Execute(w, struct {
		PageData
		Records []Produktas
	}{data, records}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//if _, err := db.Exec(query); err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	// Po įrašo ištrynimo peradresuojame vartotoją į pagrindinį puslapį

	// Po įrašo ištrynimo peradresuojame vartotoją į pagrindinį puslapį

}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	// Prisijungiame prie duomenų bazės
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Gauti įrašo ID, kurį reikia ištrinti
	id := r.FormValue("id")

	// SQL užklausa, kuri ištrina įrašą pagal ID
	query := fmt.Sprintf("DELETE FROM `produktai` WHERE produktai.id = '%s'", id)

	// Įvykdyti SQL užklausą
	if _, err := db.Exec(query); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Po įrašo ištrynimo peradresuojame vartotoją į pagrindinį puslapį
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func h1(w http.ResponseWriter, r *http.Request) {
	//	films := map[string][]Film{
	//		"Films": {
	//			{Title: "The Godfather", Director: "Francis Ford Coppola"},
	//			{Title: "The Godfather", Director: "Francis Ford Coppola"},
	//			{Title: "The Godfather", Director: "Francis Ford Coppola"},
	//		},
	//	}
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully connected to the database!")

	rows, err := db.Query("SELECT * FROM produktai")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	type Produktas struct {
		Pavadinimas string
		Kodas       int
		Parduotuve  string
		Plius       string
		Minus       string
		Keisti      string
		ID          int
	}
	var records []Produktas

	for rows.Next() {
		var record Produktas
		record.Minus = "gallery/circle-minus-1.png"
		record.Plius = "gallery/Circled_plus.png"
		record.Keisti = "gallery/keisti.png"
		if err := rows.Scan(&record.Kodas, &record.Pavadinimas, &record.Parduotuve, &record.ID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		records = append(records, record)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//}
	data := PageData{
		Title:        "Example Page",
		ImageURL:     "gallery/25694.png", // URL to your image
		Avatar:       "gallery/3837171.png",
		Parduotuve:   "gallery/1413908.png",
		Atsiliepimai: "gallery/787610-200.png",
		Nustatymai:   "gallery/setting.png",
		Ataskaita:    "gallery/1268.png",
		Apie:         "gallery/question.png",
		Lupa:         "gallery/lupa.png",
		Rodykle:      "gallery/rodykle1.png",
		Ikona:        "gallery/ikona.png",
		Bruksniai:    "gallery/bruksniai.png",
		Diagramos:    "gallery/diagramos.bmp",
		Prekes:       "gallery/prekes.png",
		Plius:        "gallery/Circled_plus.png",
		Minus:        "gallery/circle-minus-1.png",
		Keisti:       "gallery/keisti.png",
	}

	//tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl := template.Must(template.ParseFiles("duomenųBazė.html"))

	if err := tmpl.Execute(w, struct {
		PageData
		Records []Produktas
	}{data, records}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//	tmpl.Execute(w, films)

}

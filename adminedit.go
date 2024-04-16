package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	// Atsidaro ir analizuo HTML šabloną

	tmpl, err := template.ParseFiles("adminedit.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Duomenys sesijos kintamuosiuose
	data := struct {
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
		LoggedIn     bool
		Kazkas       int
		Elpastas     string
		Slaptazodis  string
		Druska       string
		Vardas       string
	}{
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
		// Čia pateikite tikrus duomenis iš sesijos
		LoggedIn:    session.Values["LoggedIn"].(bool),
		Kazkas:      session.Values["userid"].(int),
		Elpastas:    session.Values["elpastas"].(string),
		Slaptazodis: session.Values["slaptazodis"].(string),
		Druska:      session.Values["Druska"].(string),
		Vardas:      session.Values["vardas"].(string),
	}

	// Išvedame HTML puslapį su sesijos duomenimis
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func admindata(w http.ResponseWriter, r *http.Request) {
	username := "admin"
	password := "admin"
	host := "127.0.0.1"
	port := "3306"
	dbName := "pvp_admin"

	// Create the Data Source Name (DSN) string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbName)

	// Open a connection to the MySQL database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}
	defer db.Close()

	// Ping the database to check if the connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	r.ParseForm()
	id := r.FormValue("id")
	elpastas := r.FormValue("elpastas")
	vardas := r.FormValue("vardas")
	password = r.FormValue("password")

	// Atlikite užklausą
	var druska1 string
	err = db.QueryRow("SELECT druska FROM paskyra WHERE id = ?", id).Scan(&druska1)
	if err != nil {
		panic(err.Error()) // apdorokite klaidas pagal savo programos logiką
	}
	hashedPassword := hashPassword(password, druska1)
	naudotojasQuery := `
		UPDATE paskyra 
		SET  
		 	elpastas=?, 
			pilnas_vardas = ?, 
			slaptazodis = ?
		WHERE id = ?
	`

	_, err = db.Exec(naudotojasQuery,
		elpastas,
		vardas,
		hashedPassword,
		id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Gauti esamą sesiją
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Nustatyti naujas sesijos kintamąsias

	session.Values["elpastas"] = elpastas
	session.Values["slaptazodis"] = password
	session.Values["vardas"] = vardas

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Produktas struct {
	ID                 int
	Parduotuve_ID      int
	Parduotuve_adresas string
	Kodas              int
	Pavadinimas        string
	Kaina              float64
	Kategorija         string
	Sudetis            string
	Maistingumas       string
	Pagaminimo_data    string
	Galiojimo_pabaiga  string
	Parduotuve         string
	Gamintosjas        string
	Parduotuve_pav     string
	Plius              string
	Minus              string
	Keisti             string
	Gamintojas_pav     string
	Gamintojas_ID      int
	Gamintojas_salis   string
}

type Film struct {
	Title    string
	Director string
	Efektas  string
}

// EDIT BUTTON SQL CODE AND ALL FUNCTIONS
func editHandler(w http.ResponseWriter, r *http.Request) {
	username := "root"
	password := ""
	host := "127.0.0.1"
	port := "3306"
	dbName := "pvp_naudotojams"

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

	fmt.Println("Connected to MySQL database!")

	// Get the 'id' value from the request
	id := r.FormValue("id")

	// Formulate the SQL query
	query := fmt.Sprintf(`
    SELECT 
        produktas.*, 
        parduotuve.*, 
        gamintojas.*
    FROM 
        produktas 
    JOIN 
        parduotuve ON produktas.fk_parduotuve_id = parduotuve.id
    JOIN 
        gamintojas ON produktas.fk_gamintojas_id = gamintojas.id
    WHERE 
        produktas.id = '%s'
`, id)

	// Execute the SQL query to fetch the product details
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Initialize a slice to hold the product records
	var records []Produktas

	// Iterate over the rows and scan the data into Produktas struct
	for rows.Next() {
		var record Produktas

		// Scan the row into the struct fields
		if err := rows.Scan(&record.ID, &record.Kodas, &record.Pavadinimas, &record.Kaina, &record.Kategorija, &record.Sudetis, &record.Maistingumas, &record.Pagaminimo_data, &record.Galiojimo_pabaiga,
			&record.Parduotuve, &record.Gamintosjas, &record.Parduotuve_ID, &record.Parduotuve_pav, &record.Parduotuve_adresas, &record.Gamintojas_ID, &record.Gamintojas_pav, &record.Gamintojas_salis); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Append the record to the slice
		records = append(records, record)
	}

	// Check for any errors encountered during iteration
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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

	if err := tmpl.Execute(w, struct {
		PageData
		Records []Produktas
	}{data, records}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// DELETE BUTTON CODE
func deleteHandler(w http.ResponseWriter, r *http.Request) {
	username := "root"
	password := ""
	host := "127.0.0.1"
	port := "3306"
	dbName := "pvp_naudotojams"

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

	fmt.Println("Connected to MySQL database!")

	// Ping the database to check if the connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	// Gauti įrašo ID, kurį reikia ištrinti
	id := r.FormValue("id")

	// SQL užklausa, kuri ištrina įrašą pagal ID
	query := fmt.Sprintf("DELETE FROM `produktas` WHERE produktas.id = '%s'", id)

	// Įvykdyti SQL užklausą
	if _, err := db.Exec(query); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Po įrašo ištrynimo peradresuojame vartotoją į pagrindinį puslapį
	http.Redirect(w, r, "/pagrindinis", http.StatusSeeOther)
}

// AFTER LOGIN PAGE CODE "/pagrininis"
func h1(w http.ResponseWriter, r *http.Request) {
	username := "root"
	password := ""
	host := "127.0.0.1"
	port := "3306"
	dbName := "pvp_naudotojams"

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

	fmt.Println("Connected to MySQL database!")

	// Ping the database to check if the connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	//productQuery := `
	//SELECT produktas.* , parduotuve.*
	//FROM produktas
	//JOIN parduotuve ON produktas.fk_parduotuve_id = parduotuve.id`
	productQuery := fmt.Sprintf(`
    SELECT 
        produktas.*, 
        parduotuve.*, 
        gamintojas.*
    FROM 
        produktas 
    JOIN 
        parduotuve ON produktas.fk_parduotuve_id = parduotuve.id
    JOIN 
        gamintojas ON produktas.fk_gamintojas_id = gamintojas.id
`)

	//record
	// SQL užklausos rezultatai produktams
	productRows, err := db.Query(productQuery)
	if err != nil {
		log.Fatalf("Klaida vykdant užklausą: %v", err)
	}
	defer productRows.Close()

	// Čia apdoroti produktų rezultatus
	var records []Produktas
	for productRows.Next() {
		var record Produktas
		record.Minus = "gallery/circle-minus-1.png"
		record.Plius = "gallery/Circled_plus.png"
		record.Keisti = "gallery/keisti.png"
		record.Minus = "gallery/circle-minus-1.png"
		record.Plius = "gallery/Circled_plus.png"
		record.Keisti = "gallery/keisti.png"
		if err := productRows.Scan(&record.ID, &record.Kodas, &record.Pavadinimas, &record.Kaina, &record.Kategorija, &record.Sudetis, &record.Maistingumas, &record.Pagaminimo_data, &record.Galiojimo_pabaiga,
			&record.Parduotuve, &record.Gamintosjas, &record.Parduotuve_ID, &record.Parduotuve_pav, &record.Parduotuve_adresas, &record.Gamintojas_ID, &record.Gamintojas_pav, &record.Gamintojas_salis); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		records = append(records, record)
	}
	if err := productRows.Err(); err != nil {
		log.Fatalf("Klaida nuskaitant produktų rezultatus: %v", err)
		return
	}
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

}

// HTML CODE FOR INSERT PAGE PRINTINING
func h2(w http.ResponseWriter, r *http.Request) {
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
	var records []Produktas

	var record Produktas
	record.Minus = "gallery/circle-minus-1.png"
	record.Plius = "gallery/Circled_plus.png"
	record.Keisti = "gallery/keisti.png"
	record.Minus = "gallery/circle-minus-1.png"
	record.Plius = "gallery/Circled_plus.png"
	record.Keisti = "gallery/keisti.png"

	records = append(records, record)

	tmpl := template.Must(template.ParseFiles("prideti.html"))
	if err := tmpl.Execute(w, struct {
		PageData
		Records []Produktas
	}{data, records}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// INSER CODE
func insertHandler(w http.ResponseWriter, r *http.Request) {
	username := "root"
	password := ""
	host := "127.0.0.1"
	port := "3306"
	dbName := "pvp_naudotojams"

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

	fmt.Println("Connected to MySQL database!")

	// Ping the database to check if the connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Čia gauname duomenis iš formos
	gamintojasPavadinimas := r.FormValue("gamintojas_pavadinimas")
	gamintojasKilmesSalis := r.FormValue("gamintojas_kilmes_salis")
	parduotuvePavadinimas := r.FormValue("parduotuve_pavadinimas")
	parduotuveAdresas := r.FormValue("parduotuve_adresas")
	produktoBruksninisKodas := r.FormValue("produktas_bruksninis_kodas")
	produktoPavadinimas := r.FormValue("produktas_pavadinimas")
	produktoKaina := r.FormValue("produktas_kaina")
	produktoKategorija := r.FormValue("produktas_kategorija")
	produktoSudetis := r.FormValue("produktas_sudetis")
	produktoMaistingumas := r.FormValue("produktas_maistingumas")
	produktoPagaminimoData := r.FormValue("produktas_pagaminimo_data")
	produktoGaliojimoPabaigosData := r.FormValue("produktas_galiojimo_pabaigos_data")

	// Čia įrašykite savo SQL įterpimo užklausas
	query := `
			   INSERT INTO gamintojas (pavadinimas, kilmes_salis)
			   VALUES (?, ?)
		   `
	_, err = db.Exec(query, gamintojasPavadinimas, gamintojasKilmesSalis)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	query1 := `
  				 SELECT id 
   				FROM gamintojas 
   				WHERE pavadinimas = ? AND kilmes_salis = ?`
	row1 := db.QueryRow(query1, gamintojasPavadinimas, gamintojasKilmesSalis)
	var id_gamintojas int
	err = row1.Scan(&id_gamintojas)
	if err != nil {
		if err == sql.ErrNoRows {
			// Handle case where no rows were returned
		} else {
			// Handle other errors
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	query = `
				INSERT INTO parduotuve (pavadinimas, adresas)
				VALUES (?, ?)
			`
	_, err = db.Exec(query, parduotuvePavadinimas, parduotuveAdresas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	query2 := `
  				 SELECT id 
   				FROM parduotuve 
   				WHERE pavadinimas = ? AND adresas = ?`
	row2 := db.QueryRow(query2, parduotuvePavadinimas, parduotuveAdresas)
	var id_padruotuve int
	err = row2.Scan(&id_padruotuve)
	if err != nil {
		if err == sql.ErrNoRows {
			// Handle case where no rows were returned
		} else {
			// Handle other errors
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Čia įrašykite kodą, kuris įterpia duomenis į produktą
	query = `
				INSERT INTO produktas (bruksninis_kodas, pavadinimas, kaina, kategorija, sudetis, maistingumas, pagaminimo_data, galiojimo_pabaigos_data, fk_parduotuve_id, fk_gamintojas_id)
				VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
			`
	_, err = db.Exec(query, produktoBruksninisKodas, produktoPavadinimas, produktoKaina, produktoKategorija, produktoSudetis, produktoMaistingumas, produktoPagaminimoData, produktoGaliojimoPabaigosData, id_padruotuve, id_gamintojas) // 1 yra įstatyta reikšmė fk_parduotuve_id ir fk_gamintojas_id, pakeiskite ją pagal savo poreikius
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/insert", http.StatusSeeOther)

}

// SEARCH CODE
func searchHandler(w http.ResponseWriter, r *http.Request) {
	username := "root"
	password := ""
	host := "127.0.0.1"
	port := "3306"
	dbName := "pvp_naudotojams"

	// Sukurkite duomenų šaltinio vardą (DSN) eilutę
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbName)

	// Atidarome ryšį su MySQL duomenų baze
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Klaida atidarant duomenų bazės ryšį: %v", err)
	}
	defer db.Close()

	// Patikriname duomenų bazės ryšį
	err = db.Ping()
	if err != nil {
		log.Fatalf("Klaida jungiantis prie duomenų bazės: %v", err)
	}

	// Gauname įvestas reikšmes iš formos
	produkto_pavadinimas := r.FormValue("produkto_pavadinimas")
	produkto_kodas := r.FormValue("produkto_kodas")
	parduotuves_pavadinimas := r.FormValue("parduotuves_pavadinimas")

	// SQL užklausa, skirta rasti produktus, kurių pavadinimas prasideda su įvestu tekstu
	/*productQuery := `
	  SELECT *
	  FROM produktas
	  WHERE pavadinimas LIKE ? AND bruksninis_kodas LIKE ?`*/

	productQuery := `
    SELECT produktas.*, parduotuve.* 
	FROM produktas 
	JOIN parduotuve ON produktas.fk_parduotuve_id = parduotuve.id
	WHERE produktas.pavadinimas LIKE ? AND produktas.bruksninis_kodas LIKE ? AND parduotuve.pavadinimas LIKE ?`

	//record
	// SQL užklausos rezultatai produktams
	productRows, err := db.Query(productQuery, produkto_pavadinimas+"%", produkto_kodas+"%", parduotuves_pavadinimas+"%")
	if err != nil {
		log.Fatalf("Klaida vykdant užklausą: %v", err)
	}
	defer productRows.Close()

	// Čia apdoroti produktų rezultatus
	var records []Produktas
	for productRows.Next() {
		var record Produktas
		record.Minus = "gallery/circle-minus-1.png"
		record.Plius = "gallery/Circled_plus.png"
		record.Keisti = "gallery/keisti.png"
		record.Minus = "gallery/circle-minus-1.png"
		record.Plius = "gallery/Circled_plus.png"
		record.Keisti = "gallery/keisti.png"
		if err := productRows.Scan(&record.ID, &record.Kodas, &record.Pavadinimas, &record.Kaina, &record.Kategorija, &record.Sudetis, &record.Maistingumas, &record.Pagaminimo_data, &record.Galiojimo_pabaiga,
			&record.Parduotuve, &record.Gamintosjas, &record.Parduotuve_ID, &record.Parduotuve_pav, &record.Parduotuve_adresas); err != nil {
			log.Fatalf("Klaida nuskaitant produktą: %v", err)
			return
		}

		records = append(records, record)
	}
	if err := productRows.Err(); err != nil {
		log.Fatalf("Klaida nuskaitant produktų rezultatus: %v", err)
		return
	}
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
}

// HOME PAGE
func homeHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("HomePage2.html"))

	data2 := Film{
		Title:    "Example Page",
		Director: " Example page", // URL to your image
		Efektas:  "gallery/efektas(3).png",
	}

	tmpl.Execute(w, data2)

}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	username := "root"
	password := ""
	host := "127.0.0.1"
	port := "3306"
	dbName := "pvp_naudotojams"

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

	fmt.Println("Connected to MySQL database!")

	id := r.FormValue("id")
	// Gauname formos laukų duomenis
	//produktoID := r.FormValue("produkto_id")
	produktoBruksninisKodas := r.FormValue("produktas_bruksninis_kodas")
	produktoPavadinimas := r.FormValue("produktas_pavadinimas")
	produktoKaina := r.FormValue("produktas_kaina")
	produktoKategorija := r.FormValue("produktas_kategorija")
	produktoSudetis := r.FormValue("produktas_sudetis")
	produktoMaistingumas := r.FormValue("produktas_maistingumas")
	produktoPagaminimoData := r.FormValue("produktas_pagaminimo_data")
	produktoGaliojimoPabaigosData := r.FormValue("produktas_galiojimo_pabaigos_data")
	gamintojasID := r.FormValue("gamintojo_id")
	gamintojasPavadinimas := r.FormValue("gamintojo_pavadinimas")
	gamintojasKilmesSalis := r.FormValue("gamintojo_kilmes_salis")
	parduotuvesID := r.FormValue("parduotuves_id")
	parduotuvesPavadinimas := r.FormValue("parduotuves_pavadinimas")
	parduotuvesAdresas := r.FormValue("parduotuves_adresas")

	// Atnaujiname produktą
	produktoQuery := `
		UPDATE produktas 
		SET 
			bruksninis_kodas = ?, 
			pavadinimas = ?, 
			kaina = ?, 
			kategorija = ?, 
			sudetis = ?, 
			maistingumas = ?, 
			pagaminimo_data = ?, 
			galiojimo_pabaigos_data = ?
		WHERE id = ?
	`

	_, err = db.Exec(produktoQuery,
		produktoBruksninisKodas,
		produktoPavadinimas,
		produktoKaina,
		produktoKategorija,
		produktoSudetis,
		produktoMaistingumas,
		produktoPagaminimoData,
		produktoGaliojimoPabaigosData,
		id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Atnaujiname gamintoją
	gamintojoQuery := `
		UPDATE gamintojas 
		SET 
			pavadinimas = ?, 
			kilmes_salis = ?
		WHERE id = ?
	`

	_, err = db.Exec(gamintojoQuery,
		gamintojasPavadinimas,
		gamintojasKilmesSalis,
		gamintojasID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Atnaujiname parduotuvę
	parduotuvesQuery := `
		UPDATE parduotuve 
		SET 
			pavadinimas = ?, 
			adresas = ?
		WHERE id = ?
	`

	_, err = db.Exec(parduotuvesQuery,
		parduotuvesPavadinimas,
		parduotuvesAdresas,
		parduotuvesID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/pagrindinis", http.StatusSeeOther)
}

// BAR CODE SEARCH IN INSERT
func barCodeHandler(w http.ResponseWriter, r *http.Request) {
	username := "root"
	password := ""
	host := "127.0.0.1"
	port := "3306"
	dbName := "pvp_naudotojams"

	// Sukurkite duomenų šaltinio vardą (DSN) eilutę
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbName)

	// Atidarome ryšį su MySQL duomenų baze
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Klaida atidarant duomenų bazės ryšį: %v", err)
	}
	defer db.Close()

	// Patikriname duomenų bazės ryšį
	err = db.Ping()
	if err != nil {
		log.Fatalf("Klaida jungiantis prie duomenų bazės: %v", err)
	}

	// Gauname įvestas reikšmes iš formos
	barCode := r.FormValue("barCode")
	//produkto_pavadinimas := r.FormValue("barCode")
	//produkto_kodas := r.FormValue("produkto_kodas")
	//parduotuves_pavadinimas := r.FormValue("parduotuves_pavadinimas")

	// SQL užklausa, skirta rasti produktus, kurių pavadinimas prasideda su įvestu tekstu
	/*productQuery := `
	  SELECT *
	  FROM produktas
	  WHERE pavadinimas LIKE ? AND bruksninis_kodas LIKE ?`*/

	productQuery := `
    	SELECT produktas.*, parduotuve.*, gamintojas.*
		FROM produktas 
		JOIN parduotuve ON produktas.fk_parduotuve_id = parduotuve.id
		JOIN gamintojas ON produktas.fk_gamintojas_id =  gamintojas.id
		WHERE produktas.bruksninis_kodas = ?
		`

	//record
	// SQL užklausos rezultatai produktams
	productRows, err := db.Query(productQuery, barCode)
	if err != nil {
		log.Fatalf("Klaida vykdant užklausą: %v", err)
	}
	defer productRows.Close()

	// Čia apdoroti produktų rezultatus
	var records []Produktas
	for productRows.Next() {
		var record Produktas
		record.Minus = "gallery/circle-minus-1.png"
		record.Plius = "gallery/Circled_plus.png"
		record.Keisti = "gallery/keisti.png"
		record.Minus = "gallery/circle-minus-1.png"
		record.Plius = "gallery/Circled_plus.png"
		record.Keisti = "gallery/keisti.png"
		if err := productRows.Scan(&record.ID, &record.Kodas, &record.Pavadinimas, &record.Kaina, &record.Kategorija, &record.Sudetis, &record.Maistingumas, &record.Pagaminimo_data, &record.Galiojimo_pabaiga,
			&record.Parduotuve, &record.Gamintosjas, &record.Parduotuve_ID, &record.Parduotuve_pav, &record.Parduotuve_adresas, &record.Gamintojas_ID, &record.Gamintojas_pav, &record.Gamintojas_salis); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		records = append(records, record)
	}
	if err := productRows.Err(); err != nil {
		log.Fatalf("Klaida nuskaitant produktų rezultatus: %v", err)
		return
	}
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
	http.Redirect(w, r, "/pagrindinis", http.StatusSeeOther)
	//tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl := template.Must(template.ParseFiles("prideti.html"))

	if err := tmpl.Execute(w, struct {
		PageData
		Records []Produktas
	}{data, records}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

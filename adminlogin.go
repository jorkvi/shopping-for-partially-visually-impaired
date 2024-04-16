package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("asdfghjk"))

var session *sessions.Session

type PageVariables struct {
	Title        string
	ErrorMessage string
}
type User struct {
	id                           int
	elpastas                     string
	pilnas_vardas                string
	slaptazodis                  string
	druska                       string
	registracijos_data           string
	paskutinio_prisijungimo_data string
	busena                       string
	leidimai                     int
}

/*func main() {
	// Replace these credentials with your MySQL database credentials

	http.HandleFunc("/", HomePage)
	http.HandleFunc("/login", Login)
	//http.HandleFunc("/dashboard", Dashboard)
	http.HandleFunc("/logo.png", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "logo.png")
	})
	fmt.Println("Serveris veikia adresu: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}*/

func HomePage(w http.ResponseWriter, r *http.Request) {

	pageVariables := PageVariables{
		Title: "Prisijungimas",
	}

	renderLoginPage(w, pageVariables)
}

func renderLoginPage(w http.ResponseWriter, pageVariables PageVariables) {
	tmpl, err := template.New("index").Parse(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>{{.Title}}</title>
		<style>
			body {
				font-family: Arial, sans-serif;
				text-align: center;
				margin: 100px;
				background-color: #f2f2f2; /* Pakeičiame fono spalvą */
			}
	
			.container {
				display: flex;
				justify-content: center;
				width: 800px; /* Nustatome plotį */
				height: 600px; /* Nustatome aukštį */
				margin: 0 auto; /* Centruojame elementą horizontaliai */
			}
	
			.form-container {
				background-color: white; /* Pakeičiame formos fono spalvą */
				padding: 40px; /* Pridedame tarpą aplink formą */
				border-radius: 40px; /* Padidiname apvalinimą */
				width: 800px; /* Nustatome plotį */
				height: 800px; /* Nustatome aukštį */
			}
	
			input[type="text"], input[type="password"], button {
				margin-top: 50px;
				width: 70%; /* Nustatome plotį langų */
				padding: 15px;
				margin: 15px 0;
			}
	
			button {
				background-color: #4caf50; /* Keičiame mygtuko fono spalvą */
				color: white;
				border: none;
				border-radius: 20px;
				cursor: pointer;
			}
	
			.logo {
				width: 350px; /* Nustatome paveikslelio dydį */
				height: 350px;
				position: absolute; /* Nustatome absolutų pozicionavimą */
				top: 100px; /* Atstumas nuo viršaus */
				left: 1050px; /* Atstumas nuo kairės */
			}
	
			.error-message {
				color: red;
			}
		</style>
	</head>
	<body>
	<div class="container">
		<div class="logo-container">
			<img src="logo.png" alt="Logo" class="logo">
		</div>
		<div class="form-container">
			<h2 style="margin-top: 250px;">{{.Title}} </h2>
			<p class="error-message" style="font-size: 30px;">{{.ErrorMessage}}</p>
			<form action="/logine" method="post">
				<input type="text" name="email" placeholder="Elektroninis paštas" style="font-size: 30px;" required><br>
				<input type="password" name="password" placeholder="Slaptažodis" style="font-size: 30px;" required><br>
				<button type="submit" style="font-size: 30px;">Prisijungti</button>
			</form>
		</div>
	</div>
	</body>
	</html>
	

		`)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, pageVariables)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func getUserByEmail(email string) *User {
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

	fmt.Println("Connected to MySQL database!")

	// Ping the database to check if the connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	var user User
	rows, err := db.Query("SELECT * FROM paskyra WHERE elpastas = ?", email)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.id, &user.elpastas, &user.pilnas_vardas, &user.slaptazodis, &user.druska, &user.registracijos_data, &user.paskutinio_prisijungimo_data, &user.busena, &user.leidimai)
		if err != nil {
			panic(err.Error())
		} else {
			return &user
		}

	}

	return &user

}

func Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	fmt.Println("Įvestas el. paštas:", email)

	user := getUserByEmail(email)
	if user == nil {
		renderLoginPage(w, PageVariables{Title: "Prisijungimas", ErrorMessage: "Neteisingas el. paštas arba slaptažodis."})
		return
	}

	hashedPassword := hashPassword(password, user.druska)

	if hashedPassword != user.slaptazodis {
		renderLoginPage(w, PageVariables{Title: "Prisijungimas", ErrorMessage: "Neteisingas el. paštas arba slaptažodis."})
		return
	}
	// Gauti esamą sesiją
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Nustatyti prisijungimo būseną sesijoje
	session.Values["LoggedIn"] = true
	session.Values["userid"] = user.id
	session.Values["elpastas"] = user.elpastas
	session.Values["slaptazodis"] = password
	session.Values["Druska"] = user.druska
	session.Values["vardas"] = user.pilnas_vardas

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/pagrindinis", http.StatusSeeOther)
}

func hashPassword(password, salt string) string {
	// Pridedame druską prie slaptažodžio
	saltedPassword := password + salt

	// Sukuriame SHA-256 "hash"
	hasher := sha256.New()
	hasher.Write([]byte(saltedPassword))
	hashed := hasher.Sum(nil)

	// Konvertuojame į heksadecimales formatą
	hashedPassword := hex.EncodeToString(hashed)

	return hashedPassword
}

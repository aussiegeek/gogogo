package main

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"net/http"
	"time"
)

type Product struct {
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Available_On time.Time `json:"available_on"`
}

type App struct {
	db *sqlx.DB
}

func (app *App) productsHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Printf("/%s\n", request.URL.Path[1:])
	products := []Product{}
	err := app.db.Select(&products, "SELECT name, description, available_on FROM spree_products")
	if err != nil {
		fmt.Print(err)
	} else {
		json, err := json.Marshal(products)
		if err != nil {
			fmt.Print(err)
		} else {
			fmt.Fprint(response, string(json))
		}
	}
}

func main() {
	app := new(App)
	db, err := sqlx.Connect("postgres", "user=ryanbigg dbname=spree_development_go sslmode=disable")
	if err != nil {
		fmt.Print(err)
	} else {
		app.db = db
		fmt.Println("Starting HTTP server on port 8080")
		http.HandleFunc("/", app.productsHandler)
		http.ListenAndServe(":8080", nil)
	}
}

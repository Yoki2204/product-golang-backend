package main

import (
	"api/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	// To create the connection
	db := getConnection()

	//create router
	router := mux.NewRouter()
	router.HandleFunc("/products", getProducts(db)).Methods("GET")
	router.HandleFunc("/products/{id}", getProduct(db)).Methods("GET")
	router.HandleFunc("/products", createProduct(db)).Methods("POST")
	router.HandleFunc("/products/{id}", updateProduct(db)).Methods("PUT")

	//start server
	log.Fatal(http.ListenAndServe(":8000", jsonContentTypeMiddleware(router)))
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware created ")
		w.Header().Set("Access-Control-Allow-Origin", "*") 
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
		
	})
}
func getConnection() *sql.DB {
	fmt.Println("Connect to database")
	//connect to database
	db, err := sql.Open("postgres", "host=localhost user=postgres password=yokini22 dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to database")
	//create the table if it doesn't exist
	fmt.Println("To Create table")
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS products (id SERIAL PRIMARY KEY, name TEXT, type TEXT, description TEXT, price INT)")

	if err != nil {
		log.Fatal(err)
	}
	return db
}

// get all products
func getProducts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("To Select all the products from the table")
		rows, err := db.Query("SELECT * FROM products")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		products := []models.Product{}
		for rows.Next() {
			var product models.Product
			if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Description, &product.Price); err != nil {
				log.Fatal(err)
			}
			products = append(products, product)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(products)
	}
}

// get product by id
func getProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		fmt.Println("To select the product by id")
		var product models.Product
		err := db.QueryRow("SELECT * FROM products WHERE id = $1", id).Scan(&product.ID, &product.Name, &product.Type, &product.Description, &product.Price)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		fmt.Println("Selected Product "+ product.Name)
		json.NewEncoder(w).Encode(product)
	}
}

// create product
func createProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u models.Product
		json.NewDecoder(r.Body).Decode(&u)
		fmt.Println("To create product " + u.Name)

		err := db.QueryRow("INSERT INTO products (name, type, description,price) VALUES ($1, $2, $3,$4) RETURNING id", u.Name, u.Type, u.Description, u.Price).Scan(&u.ID)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(u)
	}
}

// update product
func updateProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u models.Product
		json.NewDecoder(r.Body).Decode(&u)

		vars := mux.Vars(r)
		id := vars["id"]
		fmt.Println("To update product " + u.Name)
		_, err := db.Exec("UPDATE products SET name = $1, type = $2, description = $3, price = $4 WHERE id = $5", u.Name, u.Type, u.Description, u.Price, id)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(u)
	}
}

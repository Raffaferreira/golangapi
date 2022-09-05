package api

import (
	"log"
	"net/http"
	"time"

	st "main/apis/test/models"

	service "main/apis/test/services"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func MainTestin() {
	_db, err := gorm.Open(sqlite.Open("./test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	Db = _db

	if err := Db.AutoMigrate(&st.Grocery{}); err != nil {
		panic(err)
	}

	product := mux.NewRouter().StrictSlash(true)

	s := product.PathPrefix("/products").Subrouter()
	s.HandleFunc("/allgroceries", service.AllGroceries)
	s.HandleFunc("/{name}", service.SingleGrocery).Methods("GET")
	s.HandleFunc("/", service.GroceriesToBuy).Methods("POST")
	s.HandleFunc("/{name}", service.UpdateGrocery).Methods("PUT")
	s.HandleFunc("/{name}", service.DeleteGrocery).Methods("DELETE")

	server := &http.Server{
		Handler: product,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}

// r.HandleFunc("/allgroceries", AllGroceries)
// r.HandleFunc("/groceries/{name}", SingleGrocery)
// r.HandleFunc("/groceries", GroceriesToBuy).Methods("POST")
// r.HandleFunc("/groceries/{name}", UpdateGrocery).Methods("PUT")
// r.HandleFunc("/groceries/{name}", DeleteGrocery).Methods("DELETE")

// log.Fatal(http.ListenAndServe(":10000", r))

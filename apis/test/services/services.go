package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	st "main/apis/test/models"
	"net/http"

	"github.com/gorilla/mux"
)

func AllGroceries(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint hit: returnAllGroceries")
	json.NewEncoder(w).Encode(st.GroceriesData)
}

func SingleGrocery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]

	for _, grocery := range st.GroceriesData {
		if grocery.Name == name {
			json.NewEncoder(w).Encode(grocery)
		}
	}
}

func GroceriesToBuy(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var grocery st.Grocery
	json.Unmarshal(reqBody, &grocery)
	st.GroceriesData = append(st.GroceriesData, grocery)

	json.NewEncoder(w).Encode(st.GroceriesData)

}

func DeleteGrocery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]

	for index, grocery := range st.GroceriesData {
		if grocery.Name == name {
			st.GroceriesData = append(st.GroceriesData[:index], st.GroceriesData[index+1:]...)
		}
	}

}

func UpdateGrocery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]

	for index, grocery := range st.GroceriesData {
		if grocery.Name == name {
			st.GroceriesData = append(st.GroceriesData[:index], st.GroceriesData[index+1:]...)

			var updateGrocery st.Grocery

			json.NewDecoder(r.Body).Decode(&updateGrocery)
			st.GroceriesData = append(st.GroceriesData, updateGrocery)
			fmt.Println("Endpoint hit: UpdateGroceries")
			json.NewEncoder(w).Encode(updateGrocery)
			return
		}
	}

}

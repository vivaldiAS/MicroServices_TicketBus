package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Route struct {
	ID        int     `json:"id"`
	Derpature string  `json:"derpature"`
	Arrival   string  `json:"arrival"`
	Harga     float64 `json:"harga"`
	Type      string  `json:"type"`
	Status    string  `json:"status"`
}

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// Middleware untuk mengatur CORS
func setCorsHeaders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func Index(w http.ResponseWriter, r *http.Request) {
	setCorsHeaders(w, r)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/service_routes")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, derpature, arrival, harga, type, status FROM routes")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var routes []Route
	for rows.Next() {
		var route Route
		if err := rows.Scan(&route.ID, &route.Derpature, &route.Arrival, &route.Harga, &route.Type, &route.Status); err != nil {
			log.Fatal(err)
		}
		routes = append(routes, route)
	}

	response := Response{
		Success: true,
		Data:    routes,
		Message: "Routes Retrieved Successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func Create(w http.ResponseWriter, r *http.Request) {
	setCorsHeaders(w, r)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var route Route
	if err := json.NewDecoder(r.Body).Decode(&route); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/service_routes")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO routes (derpature, arrival, harga, type, status) VALUES (?, ?, ?, ?, ?)",
		route.Derpature, route.Arrival, route.Harga, route.Type, route.Status)
	if err != nil {
		log.Fatal(err)
	}

	response := Response{
		Success: true,
		Data:    route,
		Message: "Route Created Successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func Update(w http.ResponseWriter, r *http.Request) {
	setCorsHeaders(w, r)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var route Route
	if err := json.NewDecoder(r.Body).Decode(&route); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/service_routes")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("UPDATE routes SET derpature = ?, arrival = ?, harga = ?, type = ?, status = ? WHERE id = ?",
		route.Derpature, route.Arrival, route.Harga, route.Type, route.Status, route.ID)
	if err != nil {
		log.Fatal(err)
	}

	response := Response{
		Success: true,
		Data:    route,
		Message: "Route Updated Successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	setCorsHeaders(w, r)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/service_routes")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM routes WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}

	response := Response{
		Success: true,
		Message: "Route Deleted Successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/routes", Index)         // GET
	http.HandleFunc("/routes/create", Create) // POST
	http.HandleFunc("/routes/update", Update) // PUT
	http.HandleFunc("/routes/delete", Delete) // DELETE

	fmt.Println("Server is running on port 8082...")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

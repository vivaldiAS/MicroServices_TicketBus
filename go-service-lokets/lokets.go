package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
)

type Loket struct {
	ID          int    `json:"id,omitempty"`
	NamaLoket   string `json:"nama_loket"`
	LokasiLoket string `json:"lokasi_loket"`
	AdminID     int    `json:"admin_id"`
	Status      string `json:"status"`
	BrandID     int    `json:"brand_id"`
}

func main() {
	// Membuat instance dari router mux atau handler yang Anda gunakan
	mux := http.NewServeMux()

	// Menambahkan handler untuk route /lokets
	mux.HandleFunc("/lokets", loketsHandler)
	// Menambahkan handler untuk route /lokets/{id}
	mux.HandleFunc("/lokets/", loketByIDHandler)

	fmt.Println("Server started at :8085")

	// Middleware CORS
	c := cors.AllowAll()

	// Handler CORS
	handler := c.Handler(mux)

	// Mulai server dengan handler CORS
	http.ListenAndServe(":8085", handler)
}

func loketsHandler(w http.ResponseWriter, r *http.Request) {
	// Set header untuk mendukung CORS
	w.Header().Set("Content-Type", "application/json")

	// Tangani metode GET untuk mendapatkan semua loket
	if r.Method == http.MethodGet {
		// Konek ke database
		db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/service_lokets")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()

		lokets, err := getAllLokets(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonResponse, err := json.Marshal(lokets)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}
}

func loketByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Set header untuk mendukung CORS
	w.Header().Set("Content-Type", "application/json")

	// Tangani metode GET untuk mendapatkan loket berdasarkan ID
	if r.Method == http.MethodGet {
		// Konek ke database
		db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/service_lokets")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()

		id := r.URL.Path[len("/lokets/"):]
		loket, err := getLoketByID(db, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := struct {
			Data    Loket  `json:"data"`
			Message string `json:"message"`
		}{
			Data:    loket,
			Message: "Data Berhasil",
		}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}
}

func getAllLokets(db *sql.DB) ([]Loket, error) {
	rows, err := db.Query("SELECT id, nama_loket, lokasi_loket, admin_id, status, brand_id FROM lokets")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lokets []Loket
	for rows.Next() {
		var loket Loket
		err := rows.Scan(&loket.ID, &loket.NamaLoket, &loket.LokasiLoket, &loket.AdminID, &loket.Status, &loket.BrandID)
		if err != nil {
			return nil, err
		}
		lokets = append(lokets, loket)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return lokets, nil
}

func getLoketByID(db *sql.DB, id string) (Loket, error) {
	var loket Loket
	row := db.QueryRow("SELECT id, nama_loket, lokasi_loket, admin_id, status, brand_id FROM lokets WHERE id = ?", id)
	err := row.Scan(&loket.ID, &loket.NamaLoket, &loket.LokasiLoket, &loket.AdminID, &loket.Status, &loket.BrandID)
	if err != nil {
		return Loket{}, err
	}
	return loket, nil
}

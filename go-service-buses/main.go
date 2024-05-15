package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
)

type Bus struct {
	ID            int    `json:"id"`
	Type          string `json:"type"`
	PoliceNumber  string `json:"police_number"`
	NumberOfSeats string `json:"number_of_seats"`
	NomorPintu    string `json:"nomor_pintu"`
	SupirID       int    `json:"supir_id"`
	LoketID       int    `json:"loket_id"`
	MerkID        int    `json:"merk_id"`
	Status        int    `json:"status"`
	DriverName    string `json:"driver_name"`
	DriverEmail   string `json:"driver_email"`
	NamaLoket     string `json:"nama_loket"`
}

type Response struct {
	Message string              `json:"message,omitempty"`
	Errors  map[string][]string `json:"errors,omitempty"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/bus-add", storeBusHandler)
	mux.HandleFunc("/bus-all", getAllBusesHandler)
	mux.HandleFunc("/bus/", getBusByIDHandler)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	handler := c.Handler(mux)

	fmt.Println("Bus server started at :8086")
	http.ListenAndServe(":8086", handler)
}

func storeBusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var bus Bus
	err := json.NewDecoder(r.Body).Decode(&bus)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validateBus(bus); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := saveBusToDB(bus); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := Response{
		Message: "Berhasil menambahkan bus.",
	}

	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func getAllBusesHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/service_buses")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT b.id, b.type, b.police_number, b.number_of_seats, b.merk_id, b.nomor_pintu, b.supir_id, b.loket_id, b.status, 
		u.name AS driver_name, u.email AS driver_email, l.nama_loket
		FROM buses b
		JOIN service_user.users u ON b.supir_id = u.id
		JOIN service_lokets.lokets l ON b.loket_id = l.id
	`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var buses []Bus

	for rows.Next() {
		var bus Bus
		if err := rows.Scan(&bus.ID, &bus.Type, &bus.PoliceNumber, &bus.NumberOfSeats, &bus.MerkID, &bus.NomorPintu, &bus.SupirID, &bus.LoketID, &bus.Status, &bus.DriverName, &bus.DriverEmail, &bus.NamaLoket); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		buses = append(buses, bus)
	}

	response := struct {
		Data []Bus `json:"data"`
	}{
		Data: buses,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func getBusByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	busID := r.URL.Path[len("/bus/"):]
	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/service_buses")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	row := db.QueryRow(`
		SELECT b.id, b.type, b.police_number, b.number_of_seats, b.merk_id, b.nomor_pintu, b.supir_id, b.loket_id, b.status, 
		u.name AS driver_name, u.email AS driver_email, l.nama_loket
		FROM buses b
		JOIN service_user.users u ON b.supir_id = u.id
		JOIN service_lokets.lokets l ON b.loket_id = l.id
		WHERE b.id=?
	`, busID)

	var bus Bus

	if err := row.Scan(&bus.ID, &bus.Type, &bus.PoliceNumber, &bus.NumberOfSeats, &bus.MerkID, &bus.NomorPintu, &bus.SupirID, &bus.LoketID, &bus.Status, &bus.DriverName, &bus.DriverEmail, &bus.NamaLoket); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Bus not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		Data []Bus `json:"data"`
	}{
		Data: []Bus{bus},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func validateBus(bus Bus) error {
	// Validasi bus di sini
	return nil
}

func saveBusToDB(bus Bus) error {
	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/service_buses")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(`
		INSERT INTO buses (type, police_number, number_of_seats, nomor_pintu, supir_id, loket_id)
		VALUES (?, ?, ?, ?, ?, ?)
	`, bus.Type, bus.PoliceNumber, bus.NumberOfSeats, bus.NomorPintu, bus.SupirID, bus.LoketID)
	if err != nil {
		return err
	}

	return nil
}

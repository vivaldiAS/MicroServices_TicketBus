package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
)

// Credentials struct untuk merepresentasikan kredensial pengguna
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// User struct untuk merepresentasikan pengguna yang akan diregistrasi
type User struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	PhoneNumber     string `json:"phone_number"`
	Address         string `json:"address"`
	Gender          string `json:"gender"`
	Password        string `json:"password,omitempty"`
	ConfirmPassword string `json:"confirm_password,omitempty"`
	RoleID          int    `json:"role_id"`
}

// Response struct untuk merepresentasikan respons API
type Response struct {
	Message string              `json:"message,omitempty"`
	Errors  map[string][]string `json:"errors,omitempty"`
}

// Token struct untuk merepresentasikan JWT token
type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   int64  `json:"expires_at"`
}

func main() {
	// Membuat instance dari router mux atau handler yang Anda gunakan
	mux := http.NewServeMux()

	mux.HandleFunc("/login", loginHandler)
	mux.HandleFunc("/register", registerHandler)
	mux.HandleFunc("/register-supir", registerSupirHandler)
	mux.HandleFunc("/register-adminloket", registerAdminLoketHandler)
	mux.HandleFunc("/profile", profileHandler)
	mux.HandleFunc("/logout", logoutHandler)
	mux.HandleFunc("/user/", getUserByIDHandler)
	mux.HandleFunc("/supirs", supirAllHandler)
	mux.HandleFunc("/user/update/", updateUserHandler)
	mux.HandleFunc("/admin-lokets", adminLoketAllHandler)

	// Konfigurasi CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},                            // Izinkan semua origin
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"}, // Izinkan semua metode HTTP
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	// Gunakan middleware CORS di atas handler Anda
	handler := c.Handler(mux)

	fmt.Println("Microservice started at :8081")
	// Mulai server menggunakan handler yang sudah di-set dengan CORS
	http.ListenAndServe(":8081", handler)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Set header untuk mendukung CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS ,GET, POST, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Tangani metode OPTIONS untuk Preflight CORS
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Hanya izinkan metode POST untuk login
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse body permintaan untuk mendapatkan kredensial pengguna
	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validasi kredensial pengguna dengan database dan dapatkan role ID
	var roleID string
	var email string // Variabel untuk menyimpan email
	if !isValidUser(credentials.Email, credentials.Password, &roleID, &email) {
		response := Response{
			Message: "Login gagal",
			Errors: map[string][]string{
				"credentials": {"Email atau Password salah"},
			},
		}
		jsonResponse, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsonResponse)
		return
	}

	// Jika kredensial valid, buat access token
	token, expiresAt, err := GenerateToken(credentials.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Kirim respons sukses beserta access token, email, dan role ID
	response := struct {
		Message     string `json:"message"`
		AccessToken string `json:"access_token"`
		ExpiresAt   int64  `json:"expires_at"`
		RoleID      string `json:"role_id"`
		Email       string `json:"email"` // Tambahkan email dalam respons
	}{
		Message:     "Login berhasil",
		AccessToken: token,
		ExpiresAt:   expiresAt,
		RoleID:      roleID,
		Email:       email, // Setel email dari hasil validasi
	}

	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	// Set header untuk mendukung CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Tangani metode OPTIONS untuk Preflight CORS
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Hanya izinkan metode POST untuk registrasi
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse body permintaan untuk mendapatkan data pengguna yang akan diregistrasi
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validasi input pengguna
	if err := validateUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Simpan data pengguna ke database
	if err := saveUserToDB(user, "2"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Kirim respons sukses
	response := struct {
		Message string `json:"message"`
	}{
		Message: "Registrasi berhasil",
	}

	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func registerSupirHandler(w http.ResponseWriter, r *http.Request) {
	// Set header untuk mendukung CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Tangani metode OPTIONS untuk Preflight CORS
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Hanya izinkan metode POST untuk registrasi
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse body permintaan untuk mendapatkan data supir yang akan diregistrasi
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set password default dan role_id untuk supir
	user.Password = "123456"
	user.ConfirmPassword = "123456"

	// Validasi input supir
	if err := validateUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Simpan data supir ke database
	if err := saveUserToDB(user, "3"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Kirim respons sukses
	response := struct {
		Message string `json:"message"`
	}{
		Message: "Registrasi supir berhasil",
	}

	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func registerAdminLoketHandler(w http.ResponseWriter, r *http.Request) {
	// Set header untuk mendukung CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Tangani metode OPTIONS untuk Preflight CORS
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Hanya izinkan metode POST untuk registrasi
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse body permintaan untuk mendapatkan data admin loket yang akan diregistrasi
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set password default dan role_id untuk admin loket
	user.Password = "123456"
	user.ConfirmPassword = "123456"

	// Validasi input admin loket
	if err := validateUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Simpan data admin loket ke database
	if err := saveUserToDB(user, "4"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Kirim respons sukses
	response := struct {
		Message string `json:"message"`
	}{
		Message: "Registrasi admin loket berhasil",
	}

	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	// Set header untuk mendukung CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Tangani metode OPTIONS untuk Preflight CORS
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Hanya izinkan metode GET untuk mendapatkan profil pengguna
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Mendapatkan email pengguna dari query parameter atau header atau sesuai kebutuhan Anda
	email := r.URL.Query().Get("email")
	if email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	// Mendapatkan profil pengguna dari database berdasarkan email
	userProfile, err := getUserProfileByEmail(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Mengirim respons dengan profil pengguna
	jsonProfile, err := json.Marshal(userProfile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonProfile)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	// Set header untuk mendukung CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Tangani metode OPTIONS untuk Preflight CORS
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Cari dan hapus token dari database atau catatan logout Anda
	// Misalnya, jika menggunakan JWT, token tidak perlu disimpan di sisi server
	// Jadi, di sini Anda bisa memberikan respons berhasil tanpa melakukan tindakan tambahan

	// Kirim respons logout berhasil
	response := struct {
		Message string `json:"message"`
	}{
		Message: "Successfully logged out",
	}

	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func validateUser(user User) error {
	// Validasi kebutuhan field
	if user.Name == "" || user.Email == "" || user.PhoneNumber == "" || user.Address == "" || user.Gender == "" || user.Password == "" || user.ConfirmPassword == "" {
		return fmt.Errorf("semua field harus diisi")
	}

	// Validasi kesamaan password
	if user.Password != user.ConfirmPassword {
		return fmt.Errorf("konfirmasi password tidak sesuai dengan password")
	}

	return nil
}

func saveUserToDB(user User, roleID string) error {
	// Menghubungkan ke database
	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/service_user")
	if err != nil {
		return err
	}
	defer db.Close()

	// Insert data pengguna ke database
	_, err = db.Exec("INSERT INTO users (name, email, phone_number, address, gender, password, role_id, status) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		user.Name, user.Email, user.PhoneNumber, user.Address, user.Gender, user.Password, roleID, 1)
	if err != nil {
		return err
	}

	return nil
}

func isValidUser(email, password string, roleID *string, userEmail *string) bool {
	// Menghubungkan ke database
	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/service_user")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Query database untuk pengguna dengan email yang diberikan
	var storedPassword sql.NullString
	err = db.QueryRow("SELECT password, role_id, email FROM users WHERE email = ?", email).Scan(&storedPassword, roleID, userEmail)
	if err != nil {
		return false
	}

	// Bandingkan password yang diberikan dengan yang disimpan di database
	return storedPassword.Valid && storedPassword.String == password
}

func GenerateToken(email string) (string, int64, error) {
	// Set expiry time for token
	expirationTime := time.Now().Add(24 * time.Hour).Unix()

	// Create JWT token
	claims := &jwt.StandardClaims{
		Subject:   email,
		ExpiresAt: expirationTime,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := "secret"
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", 0, err
	}

	return tokenString, expirationTime, nil
}

func getUserProfileByEmail(email string) (interface{}, error) {
	// Menghubungkan ke database
	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/service_user")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Query database untuk mendapatkan profil pengguna berdasarkan email
	var userProfile struct {
		ID            int    `json:"id"`
		Name          string `json:"name"`
		Email         string `json:"email"`
		PhoneNumber   string `json:"phone_number"`
		Address       string `json:"address"`
		Gender        string `json:"gender"`
		Photo         string `json:"photo"`
		Status        int    `json:"status"`
		RoleID        int    `json:"role_id"`
		EmailVerified string `json:"email_verified_at"`
		Role          struct {
			ID   int    `json:"id"`
			Role string `json:"role"`
		} `json:"role"`
	}

	err = db.QueryRow("SELECT u.id, u.name, u.email, u.phone_number, u.address, u.gender, u.photo, u.status, u.role_id, u.email_verified_at, r.id, r.role FROM users u JOIN roles r ON u.role_id = r.id WHERE u.email = ?", email).
		Scan(&userProfile.ID, &userProfile.Name, &userProfile.Email, &userProfile.PhoneNumber, &userProfile.Address, &userProfile.Gender, &userProfile.Photo, &userProfile.Status, &userProfile.RoleID, &userProfile.EmailVerified, &userProfile.Role.ID, &userProfile.Role.Role)
	if err != nil {
		return nil, err
	}

	return userProfile, nil
}

func supirAllHandler(w http.ResponseWriter, r *http.Request) {
	// Set header untuk mendukung CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Tangani metode OPTIONS untuk Preflight CORS
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Hanya izinkan metode GET untuk mendapatkan daftar supir
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Panggil fungsi untuk mendapatkan daftar supir dari database
	supirs, err := getSupirAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Kirim respons dengan daftar supir
	jsonSupirs, err := json.Marshal(supirs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonSupirs)
}

func getSupirAll() ([]interface{}, error) {
	// Menghubungkan ke database
	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/service_user")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Query database untuk mendapatkan semua data supir
	rows, err := db.Query("SELECT id, name, email, phone_number, address, gender, photo, status, role_id FROM users WHERE role_id = 3")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterasi melalui hasil query dan membangun slice supir
	var supirs []interface{}
	for rows.Next() {
		var supir struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Email       string `json:"email"`
			PhoneNumber string `json:"phone_number"`
			Address     string `json:"address"`
			Gender      string `json:"gender"`
			Photo       string `json:"photo"`
			Status      string `json:"status"`
			RoleID      int    `json:"role_id"`
		}
		err := rows.Scan(
			&supir.ID,
			&supir.Name,
			&supir.Email,
			&supir.PhoneNumber,
			&supir.Address,
			&supir.Gender,
			&supir.Photo,
			&supir.Status,
			&supir.RoleID,
		)
		if err != nil {
			return nil, err
		}
		supirs = append(supirs, supir)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return supirs, nil
}

func getUserByIDHandler(w http.ResponseWriter, r *http.Request) {
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

	// Mendapatkan ID pengguna dari URL
	id := r.URL.Path[len("/user/"):]

	// Mengambil pengguna dari database berdasarkan ID
	user, err := getUserByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Mengirim respons dengan data pengguna
	jsonResponse, err := json.Marshal(map[string]interface{}{
		"data": user,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func getUserByID(id string) (interface{}, error) {
	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/service_user")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var user struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
		Address     string `json:"address"`
		Gender      string `json:"gender"`
		Photo       string `json:"photo"`
		Status      int    `json:"status"`
		RoleID      int    `json:"role_id"`
	}

	err = db.QueryRow("SELECT id, name, email, phone_number, address, gender, photo, status, role_id FROM users WHERE id = ?", id).
		Scan(&user.ID, &user.Name, &user.Email, &user.PhoneNumber, &user.Address, &user.Gender, &user.Photo, &user.Status, &user.RoleID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Set header untuk mendukung CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Tangani metode OPTIONS untuk Preflight CORS
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Hanya izinkan metode PUT untuk update user
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Mendapatkan ID pengguna dari URL
	id := r.URL.Path[len("/user/update/"):]

	// Parse body permintaan untuk mendapatkan data pengguna yang akan diupdate
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update data pengguna di database
	updatedUser, err := updateUserInDB(id, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Kirim respons sukses
	response := struct {
		Data    User   `json:"data"`
		Message string `json:"message"`
	}{
		Data:    updatedUser,
		Message: "User updated successfully.",
	}

	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func updateUserInDB(id string, user User) (User, error) {
	// Menghubungkan ke database
	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/service_user")
	if err != nil {
		return User{}, err
	}
	defer db.Close()

	// Update data pengguna di database berdasarkan ID
	_, err = db.Exec("UPDATE users SET name=?, email=?, phone_number=?, address=?, gender=? WHERE id=?",
		user.Name, user.Email, user.PhoneNumber, user.Address, user.Gender, id)
	if err != nil {
		return User{}, err
	}

	// Mendapatkan data pengguna yang baru saja diperbarui
	var updatedUser User
	err = db.QueryRow("SELECT name, email, phone_number, address, gender, role_id FROM users WHERE id=?", id).Scan(
		&updatedUser.Name, &updatedUser.Email, &updatedUser.PhoneNumber, &updatedUser.Address,
		&updatedUser.Gender, &updatedUser.RoleID)
	if err != nil {
		return User{}, err
	}

	return updatedUser, nil
}
func adminLoketAllHandler(w http.ResponseWriter, r *http.Request) {
	// Set header untuk mendukung CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Tangani metode OPTIONS untuk Preflight CORS
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Hanya izinkan metode GET untuk mendapatkan daftar admin loket
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Panggil fungsi untuk mendapatkan daftar admin loket dari database
	adminLokets, err := getAdminLoketAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Kirim respons dengan daftar admin loket
	jsonAdminLokets, err := json.Marshal(adminLokets)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonAdminLokets)
}

func getAdminLoketAll() ([]interface{}, error) {
	// Menghubungkan ke database
	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/service_user")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Query database untuk mendapatkan semua data admin loket
	rows, err := db.Query("SELECT id, name, email, phone_number, address, gender, photo, status, role_id FROM users WHERE role_id = 4")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterasi melalui hasil query dan membangun slice admin loket
	var adminLokets []interface{}
	for rows.Next() {
		var adminLoket struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Email       string `json:"email"`
			PhoneNumber string `json:"phone_number"`
			Address     string `json:"address"`
			Gender      string `json:"gender"`
			Photo       string `json:"photo"`
			Status      string `json:"status"`
			RoleID      int    `json:"role_id"`
		}
		err := rows.Scan(
			&adminLoket.ID,
			&adminLoket.Name,
			&adminLoket.Email,
			&adminLoket.PhoneNumber,
			&adminLoket.Address,
			&adminLoket.Gender,
			&adminLoket.Photo,
			&adminLoket.Status,
			&adminLoket.RoleID,
		)
		if err != nil {
			return nil, err
		}
		adminLokets = append(adminLokets, adminLoket)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return adminLokets, nil
}

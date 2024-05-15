package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID     uint
	RoleID int
}

type Bus struct {
	ID uint
}

type Route struct {
	ID uint
}

var (
	dbUser   *gorm.DB
	dbRoutes *gorm.DB
	dbBuses  *gorm.DB
)

func initDB() {
	var err error

	// Connect to service_user database
	dsnUser := "user:password@tcp(localhost:3306)/service_user?charset=utf8mb4&parseTime=True&loc=Local"
	dbUser, err = gorm.Open(mysql.Open(dsnUser), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to service_user database: %v", err)
	}

	// Connect to service_routes database
	dsnRoutes := "user:password@tcp(localhost:3306)/service_routes?charset=utf8mb4&parseTime=True&loc=Local"
	dbRoutes, err = gorm.Open(mysql.Open(dsnRoutes), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to service_routes database: %v", err)
	}

	// Connect to service_buses database
	dsnBuses := "user:password@tcp(localhost:3306)/service_buses?charset=utf8mb4&parseTime=True&loc=Local"
	dbBuses, err = gorm.Open(mysql.Open(dsnBuses), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to service_buses database: %v", err)
	}
}

func CountAll(c *gin.Context) {
	var (
		supirCount int64
		busCount   int64
		routeCount int64
	)

	// Count supir
	if err := dbUser.Model(&User{}).Where("role_id = ?", 3).Count(&supirCount).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Count buses
	if err := dbBuses.Model(&Bus{}).Count(&busCount).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Count routes
	if err := dbRoutes.Model(&Route{}).Count(&routeCount).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"supir": supirCount,
		"mobil": busCount,
		"rute":  routeCount,
	})
}

func main() {
	initDB()

	r := gin.Default()

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://127.0.0.1:8000"}, // Tambahkan semua asal yang sah
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/countAll", CountAll)
	err := r.Run(":8083")
	if err != nil {
		fmt.Printf("Failed to run server: %v\n", err)
	}
}

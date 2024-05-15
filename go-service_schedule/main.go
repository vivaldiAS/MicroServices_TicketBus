package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbSchedules *gorm.DB
	dbBuses     *gorm.DB
	dbUsers     *gorm.DB
	dbRoutes    *gorm.DB
)

type Schedule struct {
	ID      uint      `gorm:"column:id"`
	Tanggal time.Time `gorm:"column:tanggal"`
	Harga   float64   `gorm:"column:harga"`
	BusID   uint      `gorm:"column:bus_id"`
	RouteID uint      `gorm:"column:route_id"`
	Status  string    `gorm:"column:status"`
}

type Bus struct {
	ID      uint   `gorm:"column:id"`
	SupirID uint   `gorm:"column:supir_id"`
	Name    string `gorm:"column:name"`
}

type User struct {
	ID   uint   `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

type Route struct {
	ID   uint   `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

type ScheduleResponse struct {
	ScheduleID uint      `json:"schedule_id"`
	Tanggal    time.Time `json:"tanggal"`
	Harga      float64   `json:"harga"`
	Bus        Bus       `json:"bus"`
	Route      Route     `json:"route"`
	SupirName  string    `json:"supir_name"`
}

func main() {
	var err error

	dsnBuses := "username:password@tcp(hostname:port)/service_buses?charset=utf8mb4&parseTime=True&loc=Local"
	dsnSchedules := "username:password@tcp(hostname:port)/service_schedules?charset=utf8mb4&parseTime=True&loc=Local"
	dsnUsers := "username:password@tcp(hostname:port)/service_user?charset=utf8mb4&parseTime=True&loc=Local"
	dsnRoutes := "username:password@tcp(hostname:port)/service_routes?charset=utf8mb4&parseTime=True&loc=Local"

	dbSchedules, err = gorm.Open(mysql.Open(dsnSchedules), &gorm.Config{})
	if err != nil {
		panic("gagal terhubung ke basis data service_schedules")
	}
	dbBuses, err = gorm.Open(mysql.Open(dsnBuses), &gorm.Config{})
	if err != nil {
		panic("gagal terhubung ke basis data service_buses")
	}
	dbUsers, err = gorm.Open(mysql.Open(dsnUsers), &gorm.Config{})
	if err != nil {
		panic("gagal terhubung ke basis data service_user")
	}
	dbRoutes, err = gorm.Open(mysql.Open(dsnRoutes), &gorm.Config{})
	if err != nil {
		panic("gagal terhubung ke basis data service_routes")
	}

	r := gin.Default()
	r.GET("/schedules", getSchedules)
	r.Run(":8084") // Ubah port di sini
}

func getSchedules(c *gin.Context) {
	var schedules []Schedule
	dbSchedules.Where("status != ?", "complete").Find(&schedules)

	var scheduleResponses []ScheduleResponse

	for _, schedule := range schedules {
		var bus Bus
		dbBuses.Where("id = ?", schedule.BusID).First(&bus)

		var user User
		dbUsers.Where("id = ?", bus.SupirID).First(&user)

		var route Route
		dbRoutes.Where("id = ?", schedule.RouteID).First(&route)

		scheduleResponse := ScheduleResponse{
			ScheduleID: schedule.ID,
			Tanggal:    schedule.Tanggal,
			Harga:      schedule.Harga,
			Bus:        bus,
			Route:      route,
			SupirName:  user.Name,
		}

		scheduleResponses = append(scheduleResponses, scheduleResponse)
	}

	c.JSON(http.StatusOK, gin.H{"data": scheduleResponses})
}

package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User Struct
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

var Db *gorm.DB

func main() {
	// Database connection
	dsn := "root:example@tcp(127.0.0.1:3306)/gobook?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	Db = db

	// Migrate the schema
	db.AutoMigrate(&User{})

	// Start Server
	e := echo.New()

	// Routing
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// User Router
	e.POST("/users", saveUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	// Run Server
	e.Logger.Fatal(e.Start(":1323"))
}

// Create User
func saveUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	if err := Db.Create(&u).Error; err != nil {
		return c.String(http.StatusBadRequest, "Something Wrong")
	}

	return c.JSON(http.StatusCreated, u)
}

// Get User
func getUser(c echo.Context) error {
	id := c.Param("id")

	var user User
	if err := Db.First(&user, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.String(http.StatusNotFound, "Not found")
		}

		return c.String(http.StatusBadRequest, "Something bad happend")
	}

	return c.JSON(http.StatusOK, user)
}

// Update User
func updateUser(c echo.Context) error {
	id := c.Param("id")

	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	if err := Db.Where("id = ?", id).Updates(u).Error; err != nil {
		return c.String(http.StatusBadRequest, "Something bad happend")
	}

	return c.String(http.StatusOK, fmt.Sprintf("User Id %v was updated", id))
}

// Delete User
func deleteUser(c echo.Context) error {
	id := c.Param("id")

	var user User
	if err := Db.Delete(&user, id).Error; err != nil {
		return c.String(http.StatusBadRequest, "Something bad happend")
	}

	return c.String(http.StatusOK, fmt.Sprintf("User Id %v was deleted", id))
}

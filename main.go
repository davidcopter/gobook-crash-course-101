package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

type jwtCustomClaims struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
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

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routing
	// Authication Router
	e.POST("/login", login)

	// Unauthenticated route
	e.GET("/public", accessible)

	// User Router
	e.POST("/users", saveUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	// Restricted group
	r := e.Group("/private")
	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("", restricted)

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

// Login
func login(c echo.Context) error {
	// Content-Type: multipart/form-data
	name := c.FormValue("name")
	password := c.FormValue("password")

	user, err := getUserByName(name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.String(http.StatusNotFound, "Not found")
		}

		return c.String(http.StatusBadRequest, "Something bad happend")
	}

	// Throws unauthorized error
	if name != user.Name || password != user.Password {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &jwtCustomClaims{
		user.ID,
		user.Name,
		user.Email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

// Get user by name
func getUserByName(name string) (User, error) {
	var user User
	if err := Db.First(&user, "name = ?", name).Error; err != nil {
		return User{}, err
	}

	return user, nil
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

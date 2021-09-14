package main

import (
	"net/http"

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

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

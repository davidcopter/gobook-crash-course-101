package main

import (
	"gobook/internal/auth"
	"gobook/internal/post"
	"gobook/internal/user"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func main() {
	// Database connection
	dsn := "root:example@tcp(127.0.0.1:3306)/gobook?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	Db = db

	// Migrate the schema:
	// https://github.com/golang-migrate/migrate
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&post.Post{})

	// Start Server
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initial Service
	authService := auth.NewService(Db)
	userService := user.NewService(Db)
	postService := post.NewService(Db)

	// Routing
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Authication Router
	e.POST("/login", authService.Login)

	// Unauthenticated route
	e.GET("/public", accessible)

	// User Router
	e.POST("/users", userService.Create)
	e.GET("/users/:id", userService.Get)
	e.PUT("/users/:id", userService.Update)
	e.DELETE("/users/:id", userService.Delete)

	// Restricted group
	r := e.Group("/private")
	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &auth.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("", restricted)

	// Post Service
	r.POST("/posts", postService.Create)
	r.GET("/posts", postService.List)
	r.GET("/posts/:id", postService.Get)
	r.PUT("/posts/:id", postService.Update)
	r.DELETE("/posts/:id", postService.Delete)

	// Run Server
	e.Logger.Fatal(e.Start(":1323"))
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*auth.JwtCustomClaims)
	name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

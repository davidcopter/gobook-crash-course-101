package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type jwtCustomClaims struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

// Login
func login(c echo.Context) error {
	// Content-Type: multipart/form-data
	name := c.FormValue("name")
	password := c.FormValue("password")

	var user User
	if err := Db.First(&user, "name = ?", name).Error; err != nil {
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

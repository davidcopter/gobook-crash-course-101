package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

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

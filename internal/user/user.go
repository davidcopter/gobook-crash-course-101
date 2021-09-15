package user

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// User Struct
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Service struct {
	Db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{Db: db}
}

// Create User
func (s *Service) Create(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	if err := s.Db.Create(&u).Error; err != nil {
		return c.String(http.StatusBadRequest, "Something Wrong")
	}

	return c.JSON(http.StatusCreated, u)
}

// Get User
func (s *Service) Get(c echo.Context) error {
	id := c.Param("id")

	var user User
	if err := s.Db.First(&user, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.String(http.StatusNotFound, "Not found")
		}

		return c.String(http.StatusBadRequest, "Something bad happend")
	}

	return c.JSON(http.StatusOK, user)
}

// Update User
func (s *Service) Update(c echo.Context) error {
	id := c.Param("id")

	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	if err := s.Db.Where("id = ?", id).Updates(u).Error; err != nil {
		return c.String(http.StatusBadRequest, "Something bad happend")
	}

	return c.String(http.StatusOK, fmt.Sprintf("User Id %v was updated", id))
}

// Delete User
func (s *Service) Delete(c echo.Context) error {
	id := c.Param("id")

	var user User
	if err := s.Db.Delete(&user, id).Error; err != nil {
		return c.String(http.StatusBadRequest, "Something bad happend")
	}

	return c.String(http.StatusOK, fmt.Sprintf("User Id %v was deleted", id))
}

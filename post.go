package main

import (
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// `Post` belongs to `User`, `UserId` is the foreign key
type Post struct {
	gorm.Model
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId uint   `json:"user_id"`
}

func createPost(c echo.Context) error {
	p := new(Post)
	if err := c.Bind(p); err != nil {
		return err
	}

	// Get jwtCustomClaims
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)

	// Assign user id
	p.UserId = claims.ID
	if err := Db.Create(&p).Error; err != nil {
		return c.String(http.StatusInternalServerError, "Something wrong")
	}

	return c.String(http.StatusCreated, "SUCCESSED")
}

func listPost(c echo.Context) error {
	var posts []Post
	if result := Db.Find(&posts); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.String(http.StatusNotFound, "Not Found")
		}
		return c.String(http.StatusInternalServerError, "Something wrong")
	}

	return c.JSON(http.StatusOK, posts)
}

func getPost(c echo.Context) error {
	id := c.Param("id")

	var post Post
	if result := Db.First(&post, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.String(http.StatusNotFound, "Not Found")
		}
		return c.String(http.StatusInternalServerError, "Something wrong")
	}

	return c.JSON(http.StatusOK, post)
}

func updatePost(c echo.Context) error {
	id := c.Param("id")

	p := new(Post)
	if err := c.Bind(p); err != nil {
		return err
	}

	if err := Db.Model(&p).Where("id = ?", id).Updates(p).Error; err != nil {
		return c.String(http.StatusInternalServerError, "Something wrong")
	}

	return c.String(http.StatusOK, "OK")
}

func deletePost(c echo.Context) error {
	id := c.Param("id")

	if err := Db.Delete(&Post{}, id).Error; err != nil {
		return c.String(http.StatusInternalServerError, "Something wrong")
	}

	return c.String(http.StatusOK, "DELETED")
}

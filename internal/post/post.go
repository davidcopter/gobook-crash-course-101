package post

import (
	"errors"
	"gobook/internal/auth"
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

type Service struct {
	Db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{Db: db}
}

func (s *Service) Create(c echo.Context) error {
	p := new(Post)
	if err := c.Bind(p); err != nil {
		return err
	}

	// Get jwtCustomClaims
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*auth.JwtCustomClaims)

	// Assign user id
	p.UserId = claims.ID
	if err := s.Db.Create(&p).Error; err != nil {
		return c.String(http.StatusInternalServerError, "Something wrong")
	}

	return c.String(http.StatusCreated, "SUCCESSED")
}

func (s *Service) List(c echo.Context) error {
	var posts []Post
	if result := s.Db.Find(&posts); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.String(http.StatusNotFound, "Not Found")
		}
		return c.String(http.StatusInternalServerError, "Something wrong")
	}

	return c.JSON(http.StatusOK, posts)
}

func (s *Service) Get(c echo.Context) error {
	id := c.Param("id")

	var post Post
	if result := s.Db.First(&post, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.String(http.StatusNotFound, "Not Found")
		}
		return c.String(http.StatusInternalServerError, "Something wrong")
	}

	return c.JSON(http.StatusOK, post)
}

func (s *Service) Update(c echo.Context) error {
	id := c.Param("id")

	p := new(Post)
	if err := c.Bind(p); err != nil {
		return err
	}

	if err := s.Db.Model(&p).Where("id = ?", id).Updates(p).Error; err != nil {
		return c.String(http.StatusInternalServerError, "Something wrong")
	}

	return c.String(http.StatusOK, "OK")
}

func (s *Service) Delete(c echo.Context) error {
	id := c.Param("id")

	if err := s.Db.Delete(&Post{}, id).Error; err != nil {
		return c.String(http.StatusInternalServerError, "Something wrong")
	}

	return c.String(http.StatusOK, "DELETED")
}

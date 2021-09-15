package post

import (
	"errors"
	"gobook/internal/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

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
	// Binding Request
	p := new(Post)
	if err := c.Bind(p); err != nil {
		return err
	}

	// Get jwtCustomClaims utils
	user := utils.GetJWTClaims(c)

	// Assign user id
	p.UserId = user.ID

	// Create to table posts using &p struct &Post{}
	if err := s.Db.Create(&p).Error; err != nil {
		return c.String(http.StatusInternalServerError, "Something wrong")
	}

	// Response
	return c.String(http.StatusCreated, "SUCCESSED")
}

func (s *Service) List(c echo.Context) error {
	// Response struct
	// `json:"title"` => content-type: application/json
	type PostList struct {
		gorm.Model
		Title string `json:"title"`
		Body  string `json:"body"`
		Name  string
	}

	var posts []PostList
	if err := s.Db.Model(&Post{}).Select("posts.*, users.name").Joins("INNER JOIN users ON posts.user_id = users.id").Limit(10).Scan(&posts).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.String(http.StatusNotFound, "Not Found")
		}
		return c.String(http.StatusInternalServerError, "Something wrong")
	}

	return c.JSON(http.StatusOK, posts)
}

func (s *Service) Get(c echo.Context) error {
	id := c.Param("id")

	var post Post
	if err := s.Db.First(&post, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
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

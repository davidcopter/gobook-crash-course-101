# Gobook Crash Course 101
 Golang crash course project: `Gobook` like Facebook  with user-service and post-service using MySQL and JWT

[TOC]

## Basic Of Backend Development

### Client-Server

[หลักการทำงานของระบบเครือข่ายแบบ Client-Server | บริษัท สยามเน็ทเวอร์ค แอนด์ คอมพิวเตอร์ จำกัด (snc.co.th)](https://www.snc.co.th/Article/Detail/111211/หลักการทำงานของระบบเครือข่ายแบบ-Client-Server)

![หลักการทำงานของระบบเครือข่ายแบบ Client-Server | บริษัท สยามเน็ทเวอร์ค แอนด์  คอมพิวเตอร์ จำกัด](https://www.snc.co.th/upload/20862/4oXmPmOvoi.png)

### HTTP Request & Response

[AnnieCannons Intro to API slides](http://anniecannons.github.io/ac-introduction-to-apis/#/)

![มายิง Service กันดีกว่าา [Android] | by Thirawat T. | Medium](http://anniecannons.github.io/ac-introduction-to-apis/img/sparkholder-api.png)

### HTTP Request Methods

- [HTTP request methods - HTTP | MDN (mozilla.org)](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods)
- [HTTP Methods for RESTful Services (restapitutorial.com)](https://www.restapitutorial.com/lessons/httpmethods.html)

![img](https://miro.medium.com/max/1324/1*OdQfB6npJJtjMDDoRzehVw.png)

An example from https://guides.rubyonrails.org/routing.html

### HTTP Status Code

[HTTP status code - A Complete List of HTTP Status Codes (infidigit.com)](https://www.infidigit.com/blog/http-status-codes/)

![List of HTTP Status Codes](https://www.infidigit.com/wp-content/uploads/2019/12/20191227_012601_0000.png)

## Echo 💕

High performance, extensible, minimalist Go web framework [Echo - High performance, minimalist Go web framework (labstack.com)](https://echo.labstack.com/)

```bash
# Initial Module
go mod init gobook

# Install ECHO
go get github.com/labstack/echo/v4
```

### Middleware

Middleware is a function chained in the HTTP request-response cycle with access to `Echo#Context` which it uses to perform a specific action, for example, logging every request or limiting the number of requests.

- Import package

```go
import (
	"github.com/labstack/echo/v4/middleware"
)
```

- Using middleware; use middle between `echo.New()` - `e.Start()`

```go
func main() {
    e := echo.New()
    // ...
    
    // Root level middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    
    // ...
    e.Logger.Fatal(e.Start(...)))
}
```

## Gorm 💕

The fantastic ORM library for Golang [GORM - The fantastic ORM library for Golang, aims to be developer friendly.](https://gorm.io/index.html)

```bash
# Install GORM
go get -u gorm.io/gorm

# Install MySQL Driver for GORM
go get -u gorm.io/driver/mysql
```

### Auto Migrate

`db.AutoMigrate()` is automatically create table in database  

## Air

☁️ Live reload for Go apps [cosmtrek/air: ☁️ Live reload for Go apps (github.com)](https://github.com/cosmtrek/air)

```bash
# Install Air
go get -u github.com/cosmtrek/air

# Inital Setup Air in project
air init
```

## JWT

JSON Web Tokens are an open, industry standard [RFC 7519](https://tools.ietf.org/html/rfc7519) method for representing claims securely between two parties. [JSON Web Tokens - jwt.io](https://jwt.io/)

```bash
# Install Go JWT
go get -u github.com/golang-jwt/jwt
```

## Dot Env

A Go port of Ruby's dotenv library (Loads environment variables from `.env`.) [joho/godotenv: A Go port of Ruby's dotenv library (Loads environment variables from `.env`.) (github.com)](https://github.com/joho/godotenv)

```bash
# Install Lib
go get github.com/joho/godotenv
```

- Create `.env` file on root directory

```
# example .envfile; key={value} variable
S3_BUCKET=YOURS3BUCKET
SECRET_KEY=YOURSECRETKEYGOESHERE
```

## Additional Resource

- [A Tour of Go (golang.org)](https://tour.golang.org/welcome/1)https://tour.golang.org/welcome/1)
- [golang-standards/project-layout: Standard Go Project Layout (github.com)](https://github.com/golang-standards/project-layout)


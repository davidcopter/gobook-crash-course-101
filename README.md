# Gobook Crash Course 101

golang crash couse project: `Gobook` like facebook with user-service and post-service using mysql and jwt

# Setup

- `Echo` setup

```bash
mkdir myapp && cd myapp
go mod init myapp
go get github.com/labstack/echo/v4
```

- `Gorm` setup

```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

## Air (Hot Reloading)

[cosmtrek/air: ☁️ Live reload for Go apps (github.com)](https://github.com/cosmtrek/air)

```sh
go get -u github.com/cosmtrek/air
```

# Package

addition package use in this project

## JWT (JSON Web Tokens)

JSON Web Tokens are an open, industry standard [RFC 7519](https://tools.ietf.org/html/rfc7519) method for representing claims securely between two parties. [JSON Web Tokens - jwt.io](https://jwt.io/)

```bash
go get -u github.com/golang-jwt/jwt
```

## Middleware (Echo)

```sh
 go get github.com/labstack/echo/middleware

 "github.com/labstack/echo/v4/middleware"
```

## Environment Variable

```sh
go get github.com/joho/godotenv
```

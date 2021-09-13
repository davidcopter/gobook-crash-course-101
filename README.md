# Gobook Crash Course 101
 golang crash couse project: `Gobook` like facebook  with user-service and post-service using mysql and jwt

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


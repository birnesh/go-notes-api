# Simple CRUD Notes API with JWT authentication using Gin


## Project initialization commands
```
go mod init
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get -u github.com/gin-gonic/gin
go get -u golang.org/x/crypto/bcrypt
go get -u github.com/golang-jwt/jwt/v5
go get -u github.com/joho/godotenv
go get -u github.com/githubnemo/CompileDaemon
go install github.com/githubnemo/CompileDaemon
```

## env variable required

``` .env
SECRET= # JWT secret 
DB_URL=localhost user=postgres password=postgres dbname=notes port=5432 sslmode=disable
PORT=8000
```


## DB migration command

Run this to create DB tables 
```
go run migrate/migrate.go
```

## hot reload dev server

```
compiledaemon --command="./go-notes-api"
```


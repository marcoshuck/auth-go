package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/marcoshuck/auth-go/pkg/controller"
	"github.com/marcoshuck/auth-go/pkg/http"
	"github.com/marcoshuck/auth-go/pkg/model"
	"github.com/marcoshuck/auth-go/pkg/repository"
	"github.com/marcoshuck/auth-go/pkg/service"
	"os"
)

func main() {
	server := http.NewServer()
	db, err := gorm.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		panic(err)
	}

	model.Migrate(db)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	err = os.Setenv("APP_SECRET_KEY", "changeme")
	if err != nil {
		panic(err)
	}

	authService := service.NewAuthService(userService, []byte(os.Getenv("APP_SECRET_KEY")))
	authController := controller.NewAuthController(authService)
	server = http.Register(server, authController)

	err = server.Run(":3000")
	if err != nil {
		panic(err)
	}
}

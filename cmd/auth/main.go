package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/marcoshuck/auth-go/pkg/controller"
	"github.com/marcoshuck/auth-go/pkg/http"
	"github.com/marcoshuck/auth-go/pkg/model"
	"github.com/marcoshuck/auth-go/pkg/repository"
	"github.com/marcoshuck/auth-go/pkg/service"
	"github.com/marcoshuck/auth-go/pkg/validator"
	"os"
)

func main() {
	// Initialize router
	router := http.NewRouter()

	// Initialize db connection
	// TODO: Replace sqlite3 with MySQL/Postgres
	db, err := gorm.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		panic(err)
	}

	// Migrate models
	model.Migrate(db)

	// Initialize validator
	v := validator.NewValidator()

	// Initialize user layers
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, v)

	// TODO: Remove this block
	err = os.Setenv("APP_SECRET_KEY", "changeme")
	if err != nil {
		panic(err)
	}

	// Initialize auth layers
	authService := service.NewAuthService(userService, v, []byte(os.Getenv("APP_SECRET_KEY")))
	authController := controller.NewAuthController(authService)

	// Register routes
	router = http.Register(router, authController)

	// Run http server
	err = router.Run(":3000")
	if err != nil {
		panic(err)
	}
}

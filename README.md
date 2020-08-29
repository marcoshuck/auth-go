# Authentication microservice
This project is an authentication microservice written in go that follows Clean Architecture and SOLID patterns.

## Installation
```bash
$ go mod download
```

## Routes
This microservice has an Authentication Controller in charge of handling authentication requests using gin.

| Route | Method | Input | Output |
| ----- | ------ | ----- | ------ |
| /auth/register | POST | [Register](https://github.com/marcoshuck/auth-go/blob/master/pkg/dto/register.go) | [Public User](https://github.com/marcoshuck/auth-go/blob/master/pkg/dto/user.go#L22) |
| /auth/login | POST | [Login](https://github.com/marcoshuck/auth-go/blob/master/pkg/dto/login.go) | [Access Token](https://github.com/marcoshuck/auth-go/blob/master/pkg/controller/auth_controller.go#L41) |

## Usage

```go
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
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## TODO
- [ ] Refresh tokens
    - [ ] Invalidation
- [ ] Logging
- [X] Validation
- [ ] CI/CD
- [ ] Error handling
- [ ] Event subscription/publishing
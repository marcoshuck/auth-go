package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/marcoshuck/auth-go/pkg/dto"
	"github.com/marcoshuck/auth-go/pkg/service"
	"net/http"
)

// AuthController groups a set of methods to login and register into the system.
type AuthController interface {
	// Login is the handler for the login route.
	Login(ctx *gin.Context)
	// Register is the handler for the register route.
	Register(ctx *gin.Context)
}

// authController is an AuthController implementation.
type authController struct {
	// authService is the service used by the auth controller handlers.
	authService service.AuthService
}

// Login is the handler that checks if the given email and password match an active user.
// If the user exists, and the password is correct, it returns an access token.
func (a authController) Login(ctx *gin.Context) {
	var body dto.Login

	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	token, expiresAt, err := a.authService.Login(body)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token": token,
		"expires_at":   expiresAt,
	})
}

// Register is the handler that creates a new user.
func (a authController) Register(ctx *gin.Context) {
	var body dto.Register
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	user, err := a.authService.Register(body)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, user.ToPublic())
}

// NewAuthController initializes a new AuthController implementation with the given service.AuthService.
func NewAuthController(authService service.AuthService) AuthController {
	return &authController{
		authService: authService,
	}
}

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/marcoshuck/auth-go/pkg/dto"
	"github.com/marcoshuck/auth-go/pkg/service"
	"net/http"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
}

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

func NewAuthController(authService service.AuthService) AuthController {
	return &authController{
		authService: authService,
	}
}

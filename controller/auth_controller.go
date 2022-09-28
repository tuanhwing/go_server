package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"goter.com.vn/server/dto"
	"goter.com.vn/server/helper"
	"goter.com.vn/server/presenter"
	"goter.com.vn/server/service"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var input struct {
		Email    string `json:"email" form:"email" binding:"required,email"`
		Password string `json:"password" form:"password" binding:"required,min=6"`
	}
	errDTO := ctx.ShouldBind(&input)

	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(input.Email, input.Password)
	if authResult != nil {
		generatedToken := c.jwtService.GenerateToken(authResult.ID.String())
		generatedRefreshToken := c.jwtService.GenerateRefreshToken(authResult.ID.String())
		toJ := &presenter.Auth{
			Token:        generatedToken,
			RefreshToken: generatedRefreshToken,
		}
		response := helper.BuildResponse(true, "OK!", toJ)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse("Please check again your credential", "Invalid Credential", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if c.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate Email", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		id, err := c.authService.CreateUser(registerDTO)
		if err != nil {
			response := helper.BuildErrorResponse("Something went wrong", err.Error(), helper.EmptyObj{})
			ctx.JSON(http.StatusConflict, response)
		}
		var output struct {
			Id string `json:"id"`
		}
		output.Id = id.String()
		response := helper.BuildResponse(true, "OK!", output)
		ctx.JSON(http.StatusOK, response)
	}
}

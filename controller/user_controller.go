package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	language "github.com/moemoe89/go-localization"
	"goter.com.vn/server/entity"
	"goter.com.vn/server/helper"
	"goter.com.vn/server/presenter"
	"goter.com.vn/server/service"
)

type UserController interface {
	Profile(context *gin.Context)
	Logout(context *gin.Context)
	RefreshToken(context *gin.Context)
}

type userController struct {
	userService  service.UserService
	jwtService   service.JWTService
	localization language.Config
}

func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *userController) Profile(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	language := GetAcceptLanguage(ctx)

	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	if !token.Valid {
		response := helper.BuildErrorResponse(c.localization.Lookup(language, "token_invalid_or_expired"), "Invalid Credential", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	id, err := entity.StringToID(fmt.Sprintf("%v", claims["user_id"]))
	if err != nil {
		response := helper.BuildErrorResponse(c.localization.Lookup(language, "id_invalid"), "Invalid Credential", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	user, err := c.userService.Profile(id)

	if err != nil {
		response := helper.BuildErrorResponse(c.localization.Lookup(language, "something_went_wrong"), "Invalid Credential", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	toJ := &presenter.User{
		ID:   id,
		Name: user.Name,
		Phone: presenter.Phone{
			DialCode:        user.Phone.DialCode,
			PhoneNumber:     user.Phone.PhoneNumber,
			FullPhoneNumber: user.Phone.FullPhoneNumber,
		},
	}
	res := helper.BuildResponse(true, c.localization.Lookup(language, "successful"), toJ)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) Logout(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	language := GetAcceptLanguage(ctx)

	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	if !token.Valid {
		response := helper.BuildErrorResponse(c.localization.Lookup(language, "token_invalid_or_expired"), "Invalid Credential", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	_, err = entity.StringToID(fmt.Sprintf("%v", claims["user_id"]))
	if err != nil {
		response := helper.BuildErrorResponse(c.localization.Lookup(language, "id_invalid"), "Invalid Credential", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	res := helper.BuildResponse(true, c.localization.Lookup(language, "successful"), nil)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) RefreshToken(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	language := GetAcceptLanguage(ctx)

	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	if !token.Valid {
		response := helper.BuildErrorResponse(c.localization.Lookup(language, "token_invalid_or_expired"), "Invalid Credential", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	id, err := entity.StringToID(fmt.Sprintf("%v", claims["user_id"]))
	if err != nil {
		response := helper.BuildErrorResponse(c.localization.Lookup(language, "id_invalid"), "Invalid Credential", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	user, err := c.userService.Profile(id)

	if err != nil {
		response := helper.BuildErrorResponse(c.localization.Lookup(language, "something_went_wrong"), "Invalid Credential", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	generatedToken := c.jwtService.GenerateToken(user.ID.String())
	generatedRefreshToken := c.jwtService.GenerateRefreshToken(user.ID.String())
	toJ := &presenter.Auth{
		AccessToken:  generatedToken,
		RefreshToken: generatedRefreshToken,
	}

	res := helper.BuildResponse(true, c.localization.Lookup(language, "successful"), toJ)
	ctx.JSON(http.StatusOK, res)
}

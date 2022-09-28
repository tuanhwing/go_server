package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"goter.com.vn/server/entity"
	"goter.com.vn/server/helper"
	"goter.com.vn/server/presenter"
	"goter.com.vn/server/service"
)

type UserController interface {
	Profile(context *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *userController) Profile(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	if !token.Valid {
		response := helper.BuildErrorResponse("Token invalid or expired", "Invalid Credential", helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	id, err := entity.StringToID(fmt.Sprintf("%v", claims["user_id"]))
	if err != nil {
		response := helper.BuildErrorResponse("ID invalid", "Invalid Credential", helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	user, err := c.userService.Profile(id)

	if err != nil {
		response := helper.BuildErrorResponse("Something went wrong!", "Invalid Credential", helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	toJ := &presenter.User{
		ID:    id,
		Email: user.Email,
		Name:  user.Name,
	}
	res := helper.BuildResponse(true, "OK!", toJ)
	context.JSON(http.StatusOK, res)
}

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	language "github.com/moemoe89/go-localization"
	"goter.com.vn/server/helper"
	"goter.com.vn/server/presenter"
	"goter.com.vn/server/service"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	CodeVerification(ctx *gin.Context)
}

type authController struct {
	authService  service.AuthService
	jwtService   service.JWTService
	localization language.Config
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService, lang language.Config) AuthController {
	return &authController{
		authService:  authService,
		jwtService:   jwtService,
		localization: lang,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var input struct {
		DialCode string `json:"dial_code" form:"dial_code" binding:"required,min=2"`
		Phone    string `json:"phone" form:"phone" binding:"required,min=6"`
	}

	language := GetAcceptLanguage(ctx)
	errDTO := ctx.ShouldBind(&input)
	if errDTO != nil {
		response := helper.BuildErrorResponse(c.localization.Lookup(language, "failed_to_process_request"), errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	isExists, err := c.authService.IsExistPhoneNumber(input.DialCode, input.Phone)
	if err != nil {
		response := helper.BuildErrorResponse(c.localization.Lookup(language, "failed_to_process_request"), err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
		return
	}

	if !isExists {
		response := helper.BuildErrorResponse(c.localization.Lookup(language, "phone_number_does_not_exists"), "", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
		return
	}

	id, err := c.authService.SendVerifyCode(input.DialCode, input.Phone)

	if err != nil {
		response := helper.BuildErrorResponse(c.localization.Lookup(language, "something_went_wrong"), err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
		return
	}

	var output struct {
		VerificationID string `json:"verification_id"`
	}
	output.VerificationID = id.String()

	response := helper.BuildResponse(true, c.localization.Lookup(language, "successful"), output)
	ctx.JSON(http.StatusOK, response)
}

func (c *authController) Register(ctx *gin.Context) {
	var input struct {
		DialCode string `json:"dial_code" form:"dial_code" binding:"required,min=2"`
		Phone    string `json:"phone" form:"phone" binding:"required,min=6"`
	}

	language := GetAcceptLanguage(ctx)
	errDTO := ctx.ShouldBind(&input)
	if errDTO != nil {
		response := helper.BuildErrorResponse(c.localization.Lookup(language, "failed_to_process_request"), errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	isExists, err := c.authService.IsExistPhoneNumber(input.DialCode, input.Phone)
	if err != nil {
		response := helper.BuildErrorResponse(c.localization.Lookup(language, "failed_to_process_request"), err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
		return
	}

	if isExists {
		response := helper.BuildErrorResponse(c.localization.Lookup(language, "phone_number_already_exists"), "", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
		return
	}

	id, err := c.authService.SendVerifyCode(input.DialCode, input.Phone)

	if err != nil {
		response := helper.BuildErrorResponse(c.localization.Lookup(language, "something_went_wrong"), err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
		return
	}

	var output struct {
		VerificationID string `json:"verification_id"`
	}
	output.VerificationID = id.String()

	response := helper.BuildResponse(true, c.localization.Lookup(language, "successful"), output)
	ctx.JSON(http.StatusOK, response)
}

func (c *authController) CodeVerification(ctx *gin.Context) {
	var input struct {
		VerificationID string `json:"verification_id" form:"verification_id" binding:"required"`
		Code           string `json:"code" form:"code" binding:"required"`
	}

	language := GetAcceptLanguage(ctx)
	errDTO := ctx.ShouldBind(&input)

	if errDTO != nil {
		response := helper.BuildErrorResponse(c.localization.Lookup(language, "failed_to_process_request"), errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	user, err := c.authService.CodeVerification(input.VerificationID, input.Code)
	if err != nil {
		response := helper.BuildErrorResponse(c.localization.Lookup(language, "failed_to_process_request"), err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
		return
	}

	generatedToken := c.jwtService.GenerateToken(user.ID.String())
	generatedRefreshToken := c.jwtService.GenerateRefreshToken(user.ID.String())
	toJ := &presenter.Auth{
		AccessToken:  generatedToken,
		RefreshToken: generatedRefreshToken,
	}
	response := helper.BuildResponse(true, c.localization.Lookup(language, "successful"), toJ)
	ctx.JSON(http.StatusOK, response)
}

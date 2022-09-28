package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"goter.com.vn/server/dto"
	"goter.com.vn/server/entity"
	"goter.com.vn/server/helper"
	"goter.com.vn/server/presenter"
	"goter.com.vn/server/service"
)

type BookController interface {
	Create(context *gin.Context)
	FinByID(context *gin.Context)
}

type bookController struct {
	bookService service.BookService
	jwtService  service.JWTService
}

func NewBookController(bookService service.BookService, jwtService service.JWTService) BookController {
	return &bookController{
		bookService: bookService,
		jwtService:  jwtService,
	}
}

func (c *bookController) Create(context *gin.Context) {
	//Input
	var bookCreateDTO dto.BookCreateDTO
	errDTO := context.ShouldBind(&bookCreateDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

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
	bookId, err := c.bookService.Create(bookCreateDTO, id)

	if err != nil {
		response := helper.BuildErrorResponse("Something went wrong!", "Invalid Credential", helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	var output struct {
		Id string `json:"id"`
	}
	output.Id = bookId.String()
	response := helper.BuildResponse(true, "OK!", output)
	context.JSON(http.StatusOK, response)

}

func (c *bookController) FinByID(context *gin.Context) {
	id, err := entity.StringToID(context.Param("id"))
	if err != nil {
		response := helper.BuildErrorResponse("ID invalid", "Invalid Credential", helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

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

	book, err := c.bookService.FinByID(id)

	if err != nil {
		response := helper.BuildErrorResponse("Something went wrong!", "Invalid Credential", helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	toJ := &presenter.Book{
		ID:          id,
		Title:       book.Title,
		Description: book.Description,
		AuthorID:    book.AuthorID,
	}
	res := helper.BuildResponse(true, "OK!", toJ)
	context.JSON(http.StatusOK, res)
}

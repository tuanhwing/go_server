package main

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"goter.com.vn/server/config"
	"goter.com.vn/server/controller"
	"goter.com.vn/server/middleware"
	"goter.com.vn/server/repository"
	"goter.com.vn/server/service"
)

var (
	//Repository
	db             *mongo.Client             = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = *repository.NewUserRepository(db)
	bookRepository repository.BookRepository = *repository.NewBookRepository(db)

	//Service
	authService service.AuthService = service.NewAuthService(userRepository)
	jwtService  service.JWTService  = service.NewJWTService()
	userService service.UserService = service.NewUserService(userRepository)
	bookService service.BookService = service.NewBookService(bookRepository)

	//Controller
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	userController controller.UserController = controller.NewUserController(userService, jwtService)
	bookController controller.BookController = controller.NewBookController(bookService, jwtService)
)

func main() {

	defer config.CloseDatabaseConnection(db)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
	}

	bookRoutes := r.Group("api/book", middleware.AuthorizeJWT(jwtService))
	{
		bookRoutes.POST("/", bookController.Create)
		bookRoutes.GET("/:id", bookController.FinByID)
	}

	r.Run(":8080")
}

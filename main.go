//go:generate go-localize -input localizations_src -output localizations
package main

import (
	"github.com/gin-gonic/gin"
	language "github.com/moemoe89/go-localization"
	"go.mongodb.org/mongo-driver/mongo"
	"goter.com.vn/server/config"
	"goter.com.vn/server/controller"
	"goter.com.vn/server/middleware"
	"goter.com.vn/server/repository"
	"goter.com.vn/server/service"
)

var (
	//Localization
	lang language.Config = initLocalization()

	//Repository
	db             *mongo.Client             = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = *repository.NewUserRepository(db)

	//Service
	authService service.AuthService = service.NewAuthService(userRepository)
	jwtService  service.JWTService  = service.NewJWTService()
	userService service.UserService = service.NewUserService(userRepository)

	//Controller
	authController controller.AuthController = controller.NewAuthController(authService, jwtService, lang)
	userController controller.UserController = controller.NewUserController(userService, jwtService)
)

func initLocalization() language.Config {
	var lang *language.Config

	// initiate the go-localization & bind some config
	cfg := language.New()
	// json file location
	cfg.BindPath("./languages.json")
	// default language
	cfg.BindMainLocale("en")

	var err error
	lang, err = cfg.Init()
	if err != nil {
		panic(err)
	}

	return *lang
}

func main() {

	defer config.CloseDatabaseConnection(db)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
		authRoutes.POST("/codeVerification", authController.CodeVerification)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/logout", userController.Logout)
		userRoutes.GET("/refreshToken", userController.RefreshToken)
	}

	r.Run(":8080")
}

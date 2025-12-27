package main

import (
	"log"

	"github.com/Danuson17-8/corn-backend/config"
	"github.com/Danuson17-8/corn-backend/db"
	"github.com/Danuson17-8/corn-backend/handlers"
	"github.com/Danuson17-8/corn-backend/middlewares"
	"github.com/Danuson17-8/corn-backend/repositories"
	"github.com/Danuson17-8/corn-backend/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Load config
	cfg := config.NewEnvConfig()

	// Connect DB
	db.ConnectDB(cfg)
	defer db.DB.Close()

	app := fiber.New(fiber.Config{
		AppName: "Corn Corn API",
	})

	// CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://corncornn.onrender.com, http://localhost:5173",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Content-Type, Authorization",
		AllowCredentials: true,
	}))

	// Init Repository
	accountRepo := &repositories.AccountRepository{DB: db.DB}
	otpRepo := &repositories.OTPRepository{DB: db.DB}
	userRepo := &repositories.UserRepository{DB: db.DB}
	contactRepo := &repositories.ContactRepository{DB: db.DB}
	menuRepo := &repositories.MenuRepository{DB: db.DB}
	promotionRepo := &repositories.PromotionRepository{DB: db.DB}

	// Init Service
	authService := &services.AuthService{AccountRepo: accountRepo}
	otpService := &services.OTPService{Repo: otpRepo}
	profileService := &services.ProfileService{UserRepo: userRepo}
	jwtService := services.NewJWTService(cfg.JWTSecret)
	contactService := &services.ContactService{Repo: contactRepo}
	menuService := &services.MenuService{Repo: menuRepo}
	promotionService := &services.PromotionService{Repo: promotionRepo}

	// Init Handler
	authHandler := &handlers.AuthHandler{
		Auth: authService,
		OTP:  otpService,
		JWT:  jwtService,
	}
	profileHandler := &handlers.ProfileHandler{Profile: profileService}
	contactHandler := &handlers.ContactHandler{Service: contactService}
	menuHandler := &handlers.MenuHandler{Service: menuService}
	promotionHandler := &handlers.PromotionHandler{Service: promotionService}

	// Routes
	auth := app.Group("/auth")
	auth.Post("/send-otp", authHandler.SendCode)
	auth.Post("/verify-otp", authHandler.VerifyCode)
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)
	auth.Post("/logout", authHandler.Logout)

	identity := app.Group("/identity")
	identity.Get("/profile",
		middlewares.RequireAuth(jwtService),
		profileHandler.GetProfile,
	)

	menu := app.Group("/menu")
	menu.Get("/corn", menuHandler.GetMenu)

	app.Get("/promotion", promotionHandler.GetActive)
	app.Post("/contact", contactHandler.Create)

	log.Fatal(app.Listen(":" + cfg.ServerPort))

}

package infrastructure

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/tu-usuario/taskflow/Usuarios/application"
	"github.com/tu-usuario/taskflow/Usuarios/infrastructure/controllers"
	"github.com/tu-usuario/taskflow/Usuarios/infrastructure/repository"
	"github.com/tu-usuario/taskflow/Usuarios/infrastructure/routes"
	"github.com/tu-usuario/taskflow/Usuarios/infrastructure/services"
)

func SetupDependencies(r chi.Router, db *sql.DB, jwtSecret string) {
	// 1. Instanciar Adaptadores de Salida (Repositorios y Servicios)
	userRepo := repository.NewUserRepositoryMySQL(db)
	tokenManager := services.NewJWTManager(jwtSecret)
	passwordService := services.NewPasswordService()

	// 2. Instanciar Casos de Uso (Aplicación)
	registerUseCase := application.NewRegisterUserUseCase(userRepo, tokenManager, passwordService)
	loginUseCase := application.NewLoginUserUseCase(userRepo, tokenManager, passwordService)

	// 3. Instanciar Controladores (Adaptadores de Entrada)
	authController := controllers.NewAuthController(registerUseCase, loginUseCase)

	// 4. Configurar Rutas
	routes.SetupAuthRoutes(r, authController)
}

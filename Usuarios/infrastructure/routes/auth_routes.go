package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/tu-usuario/taskflow/Usuarios/infrastructure/controllers"
)

func SetupAuthRoutes(r chi.Router, authController *controllers.AuthController) {
	r.Post("/register", authController.Register)
	r.Post("/login", authController.Login)
}

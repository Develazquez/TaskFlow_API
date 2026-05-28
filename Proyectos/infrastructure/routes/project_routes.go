package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/tu-usuario/taskflow/Proyectos/infrastructure/controllers"
)

func SetupProjectRoutes(r chi.Router, projectController *controllers.ProjectController) {
	r.Post("/", projectController.Create)
	r.Get("/", projectController.GetAll)
	r.Get("/{id}", projectController.GetByID)
	r.Put("/{id}", projectController.Update)
	r.Delete("/{id}", projectController.Delete)
}

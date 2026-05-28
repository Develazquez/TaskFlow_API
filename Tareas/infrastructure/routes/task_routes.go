package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/tu-usuario/taskflow/Tareas/infrastructure/controllers"
)

func SetupTaskRoutes(r chi.Router, taskController *controllers.TaskController) {
	r.Post("/", taskController.Create)
	r.Get("/", taskController.GetAll)
	r.Get("/{id}", taskController.GetByID)
	r.Put("/{id}", taskController.Update)
	r.Delete("/{id}", taskController.Delete)
	r.Patch("/{id}/complete", taskController.ToggleComplete)
}

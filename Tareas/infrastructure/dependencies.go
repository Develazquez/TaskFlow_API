package infrastructure

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/tu-usuario/taskflow/Tareas/application"
	"github.com/tu-usuario/taskflow/Tareas/infrastructure/controllers"
	"github.com/tu-usuario/taskflow/Tareas/infrastructure/repository"
	"github.com/tu-usuario/taskflow/Tareas/infrastructure/routes"
)

func SetupDependencies(r chi.Router, db *sql.DB) {
	// 1. Instanciar Adaptadores de Salida (Repositorios)
	taskRepo := repository.NewTaskRepositoryMySQL(db)

	// 2. Instanciar Casos de Uso (Aplicación)
	createUseCase := application.NewCreateTaskUseCase(taskRepo)
	getTasksUseCase := application.NewGetTasksUseCase(taskRepo)
	getTaskByIdUseCase := application.NewGetTaskByIdUseCase(taskRepo)
	updateUseCase := application.NewUpdateTaskUseCase(taskRepo)
	deleteUseCase := application.NewDeleteTaskUseCase(taskRepo)
	toggleCompleteUseCase := application.NewToggleCompleteTaskUseCase(taskRepo)

	// 3. Instanciar Controladores (Adaptadores de Entrada)
	taskController := controllers.NewTaskController(
		createUseCase,
		getTasksUseCase,
		getTaskByIdUseCase,
		updateUseCase,
		deleteUseCase,
		toggleCompleteUseCase,
	)

	// 4. Configurar Rutas
	routes.SetupTaskRoutes(r, taskController)
}

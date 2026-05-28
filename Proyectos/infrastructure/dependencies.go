package infrastructure

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/tu-usuario/taskflow/Proyectos/application"
	"github.com/tu-usuario/taskflow/Proyectos/infrastructure/controllers"
	"github.com/tu-usuario/taskflow/Proyectos/infrastructure/repository"
	"github.com/tu-usuario/taskflow/Proyectos/infrastructure/routes"
)

func SetupDependencies(r chi.Router, db *sql.DB) {
	// 1. Instanciar Adaptadores de Salida (Repositorios)
	projectRepo := repository.NewProjectRepositoryMySQL(db)

	// 2. Instanciar Casos de Uso (Aplicación)
	createUseCase := application.NewCreateProjectUseCase(projectRepo)
	getProjectsUseCase := application.NewGetProjectsUseCase(projectRepo)
	getByIdUseCase := application.NewGetProjectByIdUseCase(projectRepo)
	updateUseCase := application.NewUpdateProjectUseCase(projectRepo)
	deleteUseCase := application.NewDeleteProjectUseCase(projectRepo)

	// 3. Instanciar Controladores (Adaptadores de Entrada)
	projectController := controllers.NewProjectController(
		createUseCase,
		getProjectsUseCase,
		getByIdUseCase,
		updateUseCase,
		deleteUseCase,
	)

	// 4. Configurar Rutas
	routes.SetupProjectRoutes(r, projectController)
}

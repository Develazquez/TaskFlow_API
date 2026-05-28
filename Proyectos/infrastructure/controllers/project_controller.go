package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/tu-usuario/taskflow/Proyectos/application"
	"github.com/tu-usuario/taskflow/Proyectos/domain/entities"
	"github.com/tu-usuario/taskflow/core"
	"github.com/tu-usuario/taskflow/middlewares"
)

type ProjectController struct {
	createUseCase      *application.CreateProjectUseCase
	getProjectsUseCase *application.GetProjectsUseCase
	getByIdUseCase     *application.GetProjectByIdUseCase
	updateUseCase      *application.UpdateProjectUseCase
	deleteUseCase      *application.DeleteProjectUseCase
}

func NewProjectController(
	createUseCase *application.CreateProjectUseCase,
	getProjectsUseCase *application.GetProjectsUseCase,
	getByIdUseCase *application.GetProjectByIdUseCase,
	updateUseCase *application.UpdateProjectUseCase,
	deleteUseCase *application.DeleteProjectUseCase,
) *ProjectController {
	return &ProjectController{
		createUseCase:      createUseCase,
		getProjectsUseCase: getProjectsUseCase,
		getByIdUseCase:     getByIdUseCase,
		updateUseCase:      updateUseCase,
		deleteUseCase:      deleteUseCase,
	}
}

func (h *ProjectController) Create(w http.ResponseWriter, r *http.Request) {
	userID, ok := middlewares.GetUserID(r.Context())
	if !ok {
		core.RespondError(w, http.StatusUnauthorized, "No autorizado")
		return
	}

	var req struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		core.RespondError(w, http.StatusBadRequest, "Entrada JSON inválida")
		return
	}

	if req.Name == "" {
		core.RespondError(w, http.StatusBadRequest, "El nombre del proyecto es obligatorio")
		return
	}

	project, err := h.createUseCase.Execute(r.Context(), userID, req.Name, req.Description)
	if err != nil {
		core.RespondError(w, core.MapErrorToHTTPStatus(err), err.Error())
		return
	}

	core.RespondJSON(w, http.StatusCreated, project)
}

func (h *ProjectController) GetAll(w http.ResponseWriter, r *http.Request) {
	userID, ok := middlewares.GetUserID(r.Context())
	if !ok {
		core.RespondError(w, http.StatusUnauthorized, "No autorizado")
		return
	}

	projects, err := h.getProjectsUseCase.Execute(r.Context(), userID)
	if err != nil {
		core.RespondError(w, core.MapErrorToHTTPStatus(err), err.Error())
		return
	}

	if projects == nil {
		projects = make([]*entities.Project, 0)
	}
	
	core.RespondJSON(w, http.StatusOK, projects)
}

func (h *ProjectController) GetByID(w http.ResponseWriter, r *http.Request) {
	userID, ok := middlewares.GetUserID(r.Context())
	if !ok {
		core.RespondError(w, http.StatusUnauthorized, "No autorizado")
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		core.RespondError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	project, err := h.getByIdUseCase.Execute(r.Context(), id, userID)
	if err != nil {
		core.RespondError(w, core.MapErrorToHTTPStatus(err), err.Error())
		return
	}

	core.RespondJSON(w, http.StatusOK, project)
}

func (h *ProjectController) Update(w http.ResponseWriter, r *http.Request) {
	userID, ok := middlewares.GetUserID(r.Context())
	if !ok {
		core.RespondError(w, http.StatusUnauthorized, "No autorizado")
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		core.RespondError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	var req struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		core.RespondError(w, http.StatusBadRequest, "Entrada JSON inválida")
		return
	}

	project, err := h.updateUseCase.Execute(r.Context(), id, userID, req.Name, req.Description)
	if err != nil {
		core.RespondError(w, core.MapErrorToHTTPStatus(err), err.Error())
		return
	}

	core.RespondJSON(w, http.StatusOK, project)
}

func (h *ProjectController) Delete(w http.ResponseWriter, r *http.Request) {
	userID, ok := middlewares.GetUserID(r.Context())
	if !ok {
		core.RespondError(w, http.StatusUnauthorized, "No autorizado")
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		core.RespondError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	if err := h.deleteUseCase.Execute(r.Context(), id, userID); err != nil {
		core.RespondError(w, core.MapErrorToHTTPStatus(err), err.Error())
		return
	}

	core.RespondJSON(w, http.StatusNoContent, nil)
}

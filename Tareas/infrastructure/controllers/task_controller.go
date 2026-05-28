package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/tu-usuario/taskflow/Tareas/application"
	"github.com/tu-usuario/taskflow/Tareas/domain/entities"
	"github.com/tu-usuario/taskflow/Tareas/domain/repository"
	"github.com/tu-usuario/taskflow/core"
	"github.com/tu-usuario/taskflow/middlewares"
)

type TaskController struct {
	createUseCase         *application.CreateTaskUseCase
	getTasksUseCase       *application.GetTasksUseCase
	getTaskByIdUseCase    *application.GetTaskByIdUseCase
	updateUseCase         *application.UpdateTaskUseCase
	deleteUseCase         *application.DeleteTaskUseCase
	toggleCompleteUseCase *application.ToggleCompleteTaskUseCase
}

func NewTaskController(
	createUseCase *application.CreateTaskUseCase,
	getTasksUseCase *application.GetTasksUseCase,
	getTaskByIdUseCase *application.GetTaskByIdUseCase,
	updateUseCase *application.UpdateTaskUseCase,
	deleteUseCase *application.DeleteTaskUseCase,
	toggleCompleteUseCase *application.ToggleCompleteTaskUseCase,
) *TaskController {
	return &TaskController{
		createUseCase:         createUseCase,
		getTasksUseCase:       getTasksUseCase,
		getTaskByIdUseCase:    getTaskByIdUseCase,
		updateUseCase:         updateUseCase,
		deleteUseCase:         deleteUseCase,
		toggleCompleteUseCase: toggleCompleteUseCase,
	}
}

func (h *TaskController) Create(w http.ResponseWriter, r *http.Request) {
	userID, ok := middlewares.GetUserID(r.Context())
	if !ok {
		core.RespondError(w, http.StatusUnauthorized, "No autorizado")
		return
	}

	var req struct {
		ProjectID   *int       `json:"project_id"`
		Title       string     `json:"title"`
		Description string     `json:"description"`
		DueDate     *time.Time `json:"due_date"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		core.RespondError(w, http.StatusBadRequest, "Entrada JSON inválida")
		return
	}

	if req.Title == "" {
		core.RespondError(w, http.StatusBadRequest, "El título de la tarea es obligatorio")
		return
	}

	task, err := h.createUseCase.Execute(r.Context(), userID, req.ProjectID, req.Title, req.Description, req.DueDate)
	if err != nil {
		core.RespondError(w, core.MapErrorToHTTPStatus(err), err.Error())
		return
	}

	core.RespondJSON(w, http.StatusCreated, task)
}

func (h *TaskController) GetAll(w http.ResponseWriter, r *http.Request) {
	userID, ok := middlewares.GetUserID(r.Context())
	if !ok {
		core.RespondError(w, http.StatusUnauthorized, "No autorizado")
		return
	}

	filters := repository.TaskFilters{}
	
	if pIDStr := r.URL.Query().Get("project_id"); pIDStr != "" {
		if pID, err := strconv.Atoi(pIDStr); err == nil {
			filters.ProjectID = &pID
		}
	}

	if compStr := r.URL.Query().Get("completed"); compStr != "" {
		comp := compStr == "true"
		filters.Completed = &comp
	}

	tasks, err := h.getTasksUseCase.Execute(r.Context(), userID, filters)
	if err != nil {
		core.RespondError(w, core.MapErrorToHTTPStatus(err), err.Error())
		return
	}

	if tasks == nil {
		tasks = make([]*entities.Task, 0)
	}

	core.RespondJSON(w, http.StatusOK, tasks)
}

func (h *TaskController) GetByID(w http.ResponseWriter, r *http.Request) {
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

	task, err := h.getTaskByIdUseCase.Execute(r.Context(), id, userID)
	if err != nil {
		core.RespondError(w, core.MapErrorToHTTPStatus(err), err.Error())
		return
	}

	core.RespondJSON(w, http.StatusOK, task)
}

func (h *TaskController) Update(w http.ResponseWriter, r *http.Request) {
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
		ProjectID   *int       `json:"project_id"`
		Title       string     `json:"title"`
		Description string     `json:"description"`
		DueDate     *time.Time `json:"due_date"`
		Completed   bool       `json:"completed"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		core.RespondError(w, http.StatusBadRequest, "Entrada JSON inválida")
		return
	}

	task, err := h.updateUseCase.Execute(r.Context(), id, userID, req.ProjectID, req.Title, req.Description, req.DueDate, req.Completed)
	if err != nil {
		core.RespondError(w, core.MapErrorToHTTPStatus(err), err.Error())
		return
	}

	core.RespondJSON(w, http.StatusOK, task)
}

func (h *TaskController) Delete(w http.ResponseWriter, r *http.Request) {
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

func (h *TaskController) ToggleComplete(w http.ResponseWriter, r *http.Request) {
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

	task, err := h.toggleCompleteUseCase.Execute(r.Context(), id, userID)
	if err != nil {
		core.RespondError(w, core.MapErrorToHTTPStatus(err), err.Error())
		return
	}

	core.RespondJSON(w, http.StatusOK, task)
}

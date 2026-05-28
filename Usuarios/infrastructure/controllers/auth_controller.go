package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/tu-usuario/taskflow/Usuarios/application"
	"github.com/tu-usuario/taskflow/core"
)

type AuthController struct {
	registerUseCase *application.RegisterUserUseCase
	loginUseCase    *application.LoginUserUseCase
}

func NewAuthController(
	registerUseCase *application.RegisterUserUseCase,
	loginUseCase *application.LoginUserUseCase,
) *AuthController {
	return &AuthController{
		registerUseCase: registerUseCase,
		loginUseCase:    loginUseCase,
	}
}

func (h *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		core.RespondError(w, http.StatusBadRequest, "Entrada JSON inválida")
		return
	}

	if req.Name == "" || req.Email == "" || req.Password == "" {
		core.RespondError(w, http.StatusBadRequest, "Nombre, email y contraseña son obligatorios")
		return
	}

	user, token, err := h.registerUseCase.Execute(r.Context(), req.Name, req.Email, req.Password)
	if err != nil {
		status := core.MapErrorToHTTPStatus(err)
		core.RespondError(w, status, err.Error())
		return
	}

	core.RespondJSON(w, http.StatusCreated, map[string]interface{}{
		"user":  user,
		"token": token,
	})
}

func (h *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		core.RespondError(w, http.StatusBadRequest, "Entrada JSON inválida")
		return
	}

	token, err := h.loginUseCase.Execute(r.Context(), req.Email, req.Password)
	if err != nil {
		status := core.MapErrorToHTTPStatus(err)
		core.RespondError(w, status, err.Error())
		return
	}

	core.RespondJSON(w, http.StatusOK, map[string]string{
		"token": token,
	})
}

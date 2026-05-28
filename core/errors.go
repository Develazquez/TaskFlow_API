package core

import "errors"

var (
	ErrUserNotFound       = errors.New("usuario no encontrado")
	ErrUserAlreadyExists  = errors.New("el usuario ya existe")
	ErrInvalidCredentials = errors.New("credenciales inválidas")
	ErrProjectNotFound    = errors.New("proyecto no encontrado")
	ErrTaskNotFound       = errors.New("tarea no encontrada")
	ErrUnauthorized       = errors.New("no autorizado")
	ErrInternalServer     = errors.New("error interno del servidor")
	ErrInvalidInput       = errors.New("entrada inválida")
)

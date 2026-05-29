package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tu-usuario/taskflow/core"
	"github.com/tu-usuario/taskflow/middlewares"

	proyectosInfra "github.com/tu-usuario/taskflow/Proyectos/infrastructure"
	tareasInfra "github.com/tu-usuario/taskflow/Tareas/infrastructure"
	usuariosInfra "github.com/tu-usuario/taskflow/Usuarios/infrastructure"
)

func main() {
	// 1. Cargar configuración global (Core)
	cfg := core.LoadConfig()

	// 2. Conectar a la base de datos MySQL (Infraestructura Compartida)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error conectando a la base de datos: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("No se pudo hacer ping a la base de datos: %v", err)
	}

	log.Println("Conectado a la base de datos correctamente")

	// 3. Inicializar Router Principal
	r := chi.NewRouter()

	// Middlewares globales
	r.Use(middlewares.CorsMiddleware) // Habilitar CORS para todo el tráfico
	r.Use(middlewares.Logger)

	// 4. Registrar Rutas y Dependencias de los Módulos
	r.Route("/api/v1", func(r chi.Router) {
		
		// Módulo: Usuarios (Rutas Públicas como /register, /login)
		usuariosInfra.SetupDependencies(r, db, cfg.JWTSecret)

		// Módulo: Proyectos y Tareas (Rutas Privadas)
		r.Group(func(r chi.Router) {
			// Middleware de autenticación para el grupo de rutas privadas
			r.Use(middlewares.AuthMiddleware(cfg.JWTSecret))

			r.Route("/projects", func(r chi.Router) {
				proyectosInfra.SetupDependencies(r, db)
			})

			r.Route("/tasks", func(r chi.Router) {
				tareasInfra.SetupDependencies(r, db)
			})
		})
	})

	// 5. Iniciar Servidor
	addr := ":" + cfg.APIPort
	log.Printf("Servidor corriendo en http://localhost%s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}

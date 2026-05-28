# TaskFlow API

API REST para la gestión de tareas y proyectos, construida en Go (Golang) 1.21+ y MySQL.

## Arquitectura

La aplicación implementa una **Arquitectura Hexagonal (Ports & Adapters)** modular por características (Bounded Contexts), inspirada en principios de Domain-Driven Design (DDD).

### Estructura de Directorios

```text
taskflow/
├── Usuarios/          # Módulo de usuarios y autenticación
│   ├── domain/        # Entidades y Puertos (Interfaces)
│   ├── application/   # Casos de Uso
│   └── infrastructure/ # Adaptadores (Controladores, Repositorios, Servicios y Rutas)
├── Proyectos/         # Módulo de gestión de proyectos
│   ├── domain/
│   ├── application/
│   └── infrastructure/
├── Tareas/            # Módulo de gestión de tareas
│   ├── domain/
│   ├── application/
│   └── infrastructure/
├── core/              # Configuración y utilidades globales
├── middlewares/       # Interceptores HTTP (Auth JWT, Logger)
├── migrations/        # Scripts SQL de base de datos
├── go.mod / go.sum    # Manejo de dependencias
├── .env               # Variables de entorno
└── main.go            # Entry Point (Composition Root)
```

## Requisitos Previos

- Go 1.21+
- Servidor MySQL 8.0+

## Instalación y Ejecución

1. Clonar el repositorio.
2. Crear la base de datos ejecutando el script `migrations/01_init.sql`.
3. Renombrar `.env.example` a `.env` y configurar las credenciales de la base de datos.
4. Descargar las dependencias:
   ```bash
   go mod tidy
   ```
5. Ejecutar la aplicación:
   ```bash
   go run main.go
   ```
   El servidor iniciará por defecto en `http://localhost:8080` (o en el puerto definido en `API_PORT`).

## 🔌 Documentación de la API (Endpoints)

La API cuenta con los siguientes endpoints agrupados por módulos. Todas las rutas privadas requieren que el token JWT sea provisto en la cabecera `Authorization: Bearer <TOKEN>`.

### 📋 Vista General de Endpoints

| Módulo | Método | Endpoint | Autenticación | Descripción |
| :--- | :--- | :--- | :--- | :--- |
| **Usuarios** | `POST` | `/api/v1/register` | 🔓 Pública | Registro de nuevo usuario + Login automático |
| | `POST` | `/api/v1/login` | 🔓 Pública | Inicio de sesión (Retorna token JWT) |
| **Proyectos**| `POST` | `/api/v1/projects` | 🔒 Requiere JWT | Crear un nuevo proyecto |
| | `GET` | `/api/v1/projects` | 🔒 Requiere JWT | Obtener todos los proyectos del usuario |
| | `GET` | `/api/v1/projects/{id}` | 🔒 Requiere JWT | Obtener detalles de un proyecto específico |
| | `PUT` | `/api/v1/projects/{id}` | 🔒 Requiere JWT | Actualizar nombre/descripción de un proyecto |
| | `DELETE`| `/api/v1/projects/{id}` | 🔒 Requiere JWT | Eliminar un proyecto permanentemente |
| **Tareas** | `POST` | `/api/v1/tasks` | 🔒 Requiere JWT | Crear una nueva tarea (general o en proyecto) |
| | `GET` | `/api/v1/tasks` | 🔒 Requiere JWT | Obtener tareas del usuario (con filtros) |
| | `GET` | `/api/v1/tasks/{id}` | 🔒 Requiere JWT | Obtener detalles de una tarea específica |
| | `PUT` | `/api/v1/tasks/{id}` | 🔒 Requiere JWT | Actualizar campos de una tarea |
| | `DELETE`| `/api/v1/tasks/{id}` | 🔒 Requiere JWT | Eliminar una tarea permanentemente |
| | `PATCH`| `/api/v1/tasks/{id}/complete`| 🔒 Requiere JWT | Alternar estado de completitud de la tarea |

---

### 🔑 Detalle de Rutas: Módulo de Usuarios

#### 🟢 Registrar Usuario
* **Ruta:** `POST /api/v1/register`
* **Acceso:** 🔓 Público
* **Descripción:** Crea un nuevo registro de usuario y retorna la entidad creada junto con el token JWT de sesión.
* **Cuerpo de la Petición (`application/json`):**
  ```json
  {
    "name": "Juan Pérez",
    "email": "juan@example.com",
    "password": "mi_clave_segura"
  }
  ```
* **Respuesta Exitosa (`201 Created`):**
  ```json
  {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "name": "Juan Pérez",
      "email": "juan@example.com",
      "created_at": "2026-05-27T03:30:00Z"
    }
  }
  ```

#### 🟢 Iniciar Sesión
* **Ruta:** `POST /api/v1/login`
* **Acceso:** 🔓 Público
* **Descripción:** Autentica credenciales y devuelve un token JWT.
* **Cuerpo de la Petición (`application/json`):**
  ```json
  {
    "email": "juan@example.com",
    "password": "mi_clave_segura"
  }
  ```
* **Respuesta Exitosa (`200 OK`):**
  ```json
  {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
  ```

---

### 📂 Detalle de Rutas: Módulo de Proyectos
> [!NOTE]
> Todas las peticiones a estas rutas deben incluir la cabecera `Authorization: Bearer <TOKEN>`.

#### 🟢 Crear Proyecto
* **Ruta:** `POST /api/v1/projects`
* **Acceso:** 🔒 Privado (JWT)
* **Descripción:** Crea un proyecto propiedad del usuario autenticado.
* **Cuerpo de la Petición (`application/json`):**
  ```json
  {
    "name": "Desarrollo TaskFlow Mobile",
    "description": "Repositorio para el desarrollo del frontend de la app."
  }
  ```
* **Respuesta Exitosa (`201 Created`):**
  ```json
  {
    "id": 1,
    "user_id": 1,
    "name": "Desarrollo TaskFlow Mobile",
    "description": "Repositorio para el desarrollo del frontend de la app.",
    "created_at": "2026-05-27T03:32:00Z",
    "updated_at": "2026-05-27T03:32:00Z"
  }
  ```

#### 🔵 Obtener Proyectos
* **Ruta:** `GET /api/v1/projects`
* **Acceso:** 🔒 Privado (JWT)
* **Descripción:** Retorna la lista de todos los proyectos creados por el usuario autenticado.
* **Respuesta Exitosa (`200 OK`):**
  ```json
  [
    {
      "id": 1,
      "user_id": 1,
      "name": "Desarrollo TaskFlow Mobile",
      "description": "Repositorio para el desarrollo del frontend de la app.",
      "created_at": "2026-05-27T03:32:00Z",
      "updated_at": "2026-05-27T03:32:00Z"
    }
  ]
  ```

#### 🔵 Obtener Proyecto por ID
* **Ruta:** `GET /api/v1/projects/{id}`
* **Acceso:** 🔒 Privado (JWT)
* **Descripción:** Recupera los detalles de un proyecto específico, validando que el usuario autenticado sea el dueño.
* **Respuesta Exitosa (`200 OK`):**
  ```json
  {
    "id": 1,
    "user_id": 1,
    "name": "Desarrollo TaskFlow Mobile",
    "description": "Repositorio para el desarrollo del frontend de la app.",
    "created_at": "2026-05-27T03:32:00Z",
    "updated_at": "2026-05-27T03:32:00Z"
  }
  ```

#### 🟡 Actualizar Proyecto
* **Ruta:** `PUT /api/v1/projects/{id}`
* **Acceso:** 🔒 Privado (JWT)
* **Descripción:** Modifica el nombre y descripción de un proyecto existente propiedad del usuario.
* **Cuerpo de la Petición (`application/json`):**
  ```json
  {
    "name": "Desarrollo TaskFlow Mobile (Fase 1)",
    "description": "Repositorio actualizado para la primera versión de la app móvil."
  }
  ```
* **Respuesta Exitosa (`200 OK`):**
  ```json
  {
    "id": 1,
    "user_id": 1,
    "name": "Desarrollo TaskFlow Mobile (Fase 1)",
    "description": "Repositorio actualizado para la primera versión de la app móvil.",
    "created_at": "2026-05-27T03:32:00Z",
    "updated_at": "2026-05-27T03:40:00Z"
  }
  ```

#### 🔴 Eliminar Proyecto
* **Ruta:** `DELETE /api/v1/projects/{id}`
* **Acceso:** 🔒 Privado (JWT)
* **Descripción:** Elimina definitivamente el proyecto especificado.
* **Respuesta Exitosa (`204 No Content`):** *(Sin cuerpo de respuesta)*

---

### 📝 Detalle de Rutas: Módulo de Tareas
> [!NOTE]
> Todas las peticiones a estas rutas deben incluir la cabecera `Authorization: Bearer <TOKEN>`.

#### 🟢 Crear Tarea
* **Ruta:** `POST /api/v1/tasks`
* **Acceso:** 🔒 Privado (JWT)
* **Descripción:** Crea una tarea. Si se proporciona `project_id`, valida que dicho proyecto pertenezca al usuario.
* **Cuerpo de la Petición (`application/json`):**
  ```json
  {
    "project_id": 1,
    "title": "Diseñar Mockups UX/UI",
    "description": "Hacer wireframes iniciales en Figma",
    "due_date": "2026-06-05T23:59:59Z"
  }
  ```
  *Nota: `project_id` puede ser `null` si es una tarea personal/libre; `due_date` es opcional.*
* **Respuesta Exitosa (`201 Created`):**
  ```json
  {
    "id": 10,
    "user_id": 1,
    "project_id": 1,
    "title": "Diseñar Mockups UX/UI",
    "description": "Hacer wireframes iniciales en Figma",
    "completed": false,
    "due_date": "2026-06-05T23:59:59Z",
    "created_at": "2026-05-27T03:45:00Z",
    "updated_at": "2026-05-27T03:45:00Z"
  }
  ```

#### 🔵 Obtener Tareas (con Filtros)
* **Ruta:** `GET /api/v1/tasks`
* **Acceso:** 🔒 Privado (JWT)
* **Descripción:** Obtiene la lista de tareas del usuario. Admite parámetros de consulta opcionales para filtrado.
* **Parámetros Query de Filtrado:**
  * `project_id`: Filtra por ID de proyecto (ej. `?project_id=1`).
  * `completed`: Filtra por estado de completitud (ej. `?completed=true` o `?completed=false`).
* **Respuesta Exitosa (`200 OK`):**
  ```json
  [
    {
      "id": 10,
      "user_id": 1,
      "project_id": 1,
      "title": "Diseñar Mockups UX/UI",
      "description": "Hacer wireframes iniciales en Figma",
      "completed": false,
      "due_date": "2026-06-05T23:59:59Z",
      "created_at": "2026-05-27T03:45:00Z",
      "updated_at": "2026-05-27T03:45:00Z"
    }
  ]
  ```

#### 🔵 Obtener Tarea por ID
* **Ruta:** `GET /api/v1/tasks/{id}`
* **Acceso:** 🔒 Privado (JWT)
* **Descripción:** Recupera la información detallada de una tarea específica del usuario.
* **Respuesta Exitosa (`200 OK`):**
  ```json
  {
    "id": 10,
    "user_id": 1,
    "project_id": 1,
    "title": "Diseñar Mockups UX/UI",
    "description": "Hacer wireframes iniciales en Figma",
    "completed": false,
    "due_date": "2026-06-05T23:59:59Z",
    "created_at": "2026-05-27T03:45:00Z",
    "updated_at": "2026-05-27T03:45:00Z"
  }
  ```

#### 🟡 Actualizar Tarea
* **Ruta:** `PUT /api/v1/tasks/{id}`
* **Acceso:** 🔒 Privado (JWT)
* **Descripción:** Actualiza de manera completa los atributos de una tarea.
* **Cuerpo de la Petición (`application/json`):**
  ```json
  {
    "project_id": 1,
    "title": "Diseñar Mockups UX/UI - Modificado",
    "description": "Mockups actualizados tras sesión de feedback.",
    "due_date": "2026-06-10T23:59:59Z",
    "completed": false
  }
  ```
* **Respuesta Exitosa (`200 OK`):**
  ```json
  {
    "id": 10,
    "user_id": 1,
    "project_id": 1,
    "title": "Diseñar Mockups UX/UI - Modificado",
    "description": "Mockups actualizados tras sesión de feedback.",
    "completed": false,
    "due_date": "2026-06-10T23:59:59Z",
    "created_at": "2026-05-27T03:45:00Z",
    "updated_at": "2026-05-27T03:50:00Z"
  }
  ```

#### 🔴 Eliminar Tarea
* **Ruta:** `DELETE /api/v1/tasks/{id}`
* **Acceso:** 🔒 Privado (JWT)
* **Descripción:** Elimina de manera permanente la tarea indicada.
* **Respuesta Exitosa (`204 No Content`):** *(Sin cuerpo de respuesta)*

#### 🟣 Alternar Estado de Completado (Toggle)
* **Ruta:** `PATCH /api/v1/tasks/{id}/complete`
* **Acceso:** 🔒 Privado (JWT)
* **Descripción:** Invierte dinámicamente el estado del campo `completed` (de `true` a `false` o viceversa).
* **Respuesta Exitosa (`200 OK`):**
  ```json
  {
    "id": 10,
    "user_id": 1,
    "project_id": 1,
    "title": "Diseñar Mockups UX/UI - Modificado",
    "description": "Mockups actualizados tras sesión de feedback.",
    "completed": true,
    "due_date": "2026-06-10T23:59:59Z",
    "created_at": "2026-05-27T03:45:00Z",
    "updated_at": "2026-05-27T03:52:00Z"
  }
  ```
# TaskFlow_API

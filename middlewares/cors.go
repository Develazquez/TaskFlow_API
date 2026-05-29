package middlewares

import "net/http"

// CorsMiddleware maneja las cabeceras CORS para permitir tráfico de cualquier origen
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Permitir cualquier origen
		w.Header().Set("Access-Control-Allow-Origin", "*")
		
		// Métodos permitidos
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		
		// Cabeceras permitidas
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
		
		// Exponer cabeceras en la respuesta si es necesario
		w.Header().Set("Access-Control-Expose-Headers", "Link")
		
		// Permitir el uso de credenciales (cookies, authorization headers, etc.)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Max-Age", "300") // Caché preflight de 5 minutos

		// Si es una petición OPTIONS (Preflight), responder con 200 OK directamente
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

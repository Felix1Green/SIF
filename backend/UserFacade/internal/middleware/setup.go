package middleware

import (
	"UserFacade/internal/middleware/cors"
	"net/http"
)

func SetupMiddleware(handler http.Handler) http.Handler {
	corsMiddleware := cors.MiddlewareCors(handler)

	return corsMiddleware
}

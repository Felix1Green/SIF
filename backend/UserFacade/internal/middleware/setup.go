package middleware

import (
	"github.com/Felix1Green/SIF/backend/UserFacade/internal/middleware/cors"
	"net/http"
)

func SetupMiddleware(handler http.Handler) http.Handler {
	corsMiddleware := cors.MiddlewareCors(handler)

	return corsMiddleware
}

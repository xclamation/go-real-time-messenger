package routes

import (
	"net/http"

	"github.com/xclamation/go-real-time-messenger/internal/handlers"
)

func InitializeRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/", handlers.HomeHandler)

	return mux
}

package statuses

import (
	"net/http"
	"yatter-backend-go/app/app"

	"github.com/go-chi/chi"
)

// Implementation of handler
type handler struct {
	app *app.App
}

// Create Handler for `/v1/statuses/`
func NewRouter(app *app.App) http.Handler {
	r := chi.NewRouter()

	h := &handler{app: app}
	r.Post("/", h.Create)
	r.Get("/{id}", h.SearchById)
	r.Delete("/{id}", h.DeleteById)

	return r
}
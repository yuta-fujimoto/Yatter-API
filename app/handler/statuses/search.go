package statuses

import (
	"encoding/json"
	"net/http"
	"strconv"
	"yatter-backend-go/app/handler/httperror"

	"github.com/go-chi/chi"
)


func (h *handler) SearchById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ids := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(ids, 10, 64)
	if err != nil {
		httperror.BadRequest(w, err)
		return
	}

	statusDomain := h.app.Dao.Status()
	status, err := statusDomain.FindById(ctx, id)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}
	if status == nil {
		httperror.Error(w, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(status); err != nil {
		httperror.InternalServerError(w, err)
	}
}
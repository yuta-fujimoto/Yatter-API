package statuses

import (
	"net/http"
	"strconv"
	"yatter-backend-go/app/handler/httperror"

	"github.com/go-chi/chi"
)

func (h *handler) DeleteById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ids := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(ids, 10, 64)
	if err != nil {
		httperror.BadRequest(w, err)
		return
	}

	statusDomain := h.app.Dao.Status()
	if err = statusDomain.Delete(ctx, id); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
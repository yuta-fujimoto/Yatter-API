package accounts

import (
	"encoding/json"
	"net/http"
	"yatter-backend-go/app/handler/httperror"

	"github.com/go-chi/chi"
)

type GetRequest struct {
	Username string
}

func (h* handler) SearchUserByName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	username := chi.URLParam(r, "username")

	accountDomain := h.app.Dao.Account()
	account, err := accountDomain.FindByUsername(ctx, username)
	if err != nil {
		httperror.InternalServerError(w, err)
		return 
	}

	if account == nil {
		httperror.Error(w, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(account); err != nil {
		httperror.InternalServerError(w, err)
		return 
	}
}
package statuses

import (
	"encoding/json"
	"net/http"

	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/httperror"
)

type AddRequest struct {
	Content string
}

// Handle request for `POST /v1/accounts`
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httperror.BadRequest(w, err)
		return
	}

	newStatus := &object.Status{
		Content: &req.Content,
		// TODO: set valid value
		AccountID: 1,
	}

	statusDomain := h.app.Dao.Status() // domain/repository の取得
	err := statusDomain.Create(ctx, newStatus)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	if err := json.NewEncoder(w).Encode(newStatus); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}

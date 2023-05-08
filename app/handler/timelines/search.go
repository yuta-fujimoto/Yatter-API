package timelines

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"
	"yatter-backend-go/app/handler/httperror"
)

const (
	LimitMax = 80
	LimitMin = 40
	LIMIT = "limit"
	SinceId = "since_id"
	MaxId = "max_id"
	OnlyMedia = "only_media"
)

func (h *handler) searchPublic(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var sinceId, maxId, limit int64
	var onlyMedia bool
	var err error
	query := r.URL.Query()
	// HACK: use Has
	if query.Get(SinceId) != "" {
		sinceId, err = strconv.ParseInt(query.Get(SinceId), 10, 64)
	} else {
		sinceId = 0
	}
	if query.Get(MaxId) != "" {
		maxId, err = strconv.ParseInt(query.Get(MaxId), 10, 64)
	} else {
		maxId = math.MaxInt64
	}
	if query.Get(LIMIT) != "" {
		limit, err = strconv.ParseInt(query.Get(LIMIT), 10, 64)
		limit = int64(math.Max(float64(limit), LimitMax))
	} else {
		limit = LimitMin
	}
	// TODO media
	if query.Get(OnlyMedia) != "" {
		onlyMedia, err = strconv.ParseBool(query.Get(OnlyMedia))
	} else {
		onlyMedia = false
	}
	if err != nil {
		httperror.BadRequest(w, err)
		return
	}

	statusRepository := h.app.Dao.Status()
	timelines, err := statusRepository.FindMany(ctx, sinceId, maxId, limit, onlyMedia)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(timelines); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}

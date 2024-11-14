package library

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func GetID(w http.ResponseWriter, r *http.Request) int {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		Response400(w, "Invalid category ID")
		return 0
	}
	return id
}

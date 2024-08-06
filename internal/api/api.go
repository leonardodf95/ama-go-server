package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/leonardodf95/ama-go-server/internal/store/pgstore"
)

type ApiHandler struct {
	q *pgstore.Queries
	r *chi.Mux
}

func (h ApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

func NewHandler(q *pgstore.Queries) http.Handler {
	a := ApiHandler{
		q: q,
	}

	r := chi.NewRouter()

	a.r = r
	return a
}

package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes(h *Handler) http.Handler {
	r := chi.NewRouter()
	r.Get("/ping", h.Ping)
	r.Post("/identities/telegram", h.IdentitiesTelegram)
	return r
}

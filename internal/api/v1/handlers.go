package v1

import "net/http"

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}
func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("pong"))
}

func (h *Handler) IdentitiesTelegram(w http.ResponseWriter, r *http.Request) {

}

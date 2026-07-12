package email

import (
	"net/http"
	"verify/3-validation-api/configs"
)

type EmailHandlerDeps struct {
	*configs.EmailConfig
}

type EmailHandler struct {
	*configs.EmailConfig
}

func NewEmailHandlers(router *http.ServeMux, config *configs.EmailConfig) {
	handler := &EmailHandler{
		EmailConfig: config,
	}
	router.HandleFunc("POST /send", handler.Send)
	router.HandleFunc("GET /verify/{hash}", handler.Verify)
}

func (h *EmailHandler) Send(w http.ResponseWriter, r *http.Request) {}

func (h *EmailHandler) Verify(w http.ResponseWriter, r *http.Request) {}

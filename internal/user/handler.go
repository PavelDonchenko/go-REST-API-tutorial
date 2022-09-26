package user

import (
	"github.com/PavelDonchenko/40projects/rest-api-tutorial/internal/handlers"
	"github.com/PavelDonchenko/40projects/rest-api-tutorial/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var _ handlers.Handler = &handler{}

type handler struct {
	logger logging.Logger
}

func NewHandler(logger logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc("GET", "/users", h.GetAllUsers)
	router.HandlerFunc("GET", "/users/:id", h.GetUserById)
	router.HandlerFunc("POST", "/users/", h.CreateUser)
	router.HandlerFunc("PUT", "/users/:id", h.UpdateUser)
	router.HandlerFunc("PATCH", "/users/:id", h.ParticalUpdateUser)
	router.HandlerFunc("DELETE", "/users/:id", h.DeleteUser)
}
func (h *handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bla bla bla"))
}

func (h *handler) GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bla bla bla"))
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bla bla bla"))
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bla bla bla"))
}

func (h *handler) ParticalUpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bla bla bla"))
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bla bla bla"))
}


package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/farrukhnaveed/go-ecom/types"
	"github.com/farrukhnaveed/go-ecom/utils"
)

type Handler struct {

}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	// router.HandleFunc("/users", h.CreateUser).Methods("POST")
	// router.HandleFunc("/users/{id}", h.GetUser).Methods("GET")
	// router.HandleFunc("/users/{id}", h.UpdateUser).Methods("PUT")
	// router.HandleFunc("/users/{id}", h.DeleteUser).Methods("DELETE")
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload RegisterUserPayload

	err := utils.ParseJSON(r, payload);
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	
}

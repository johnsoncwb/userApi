package api

import (
	"context"
	"encoding/json"
	"github.com/johnsoncwb/userApi/internal/core/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type UsersHandlers struct {
	srvc service.UserService
}

func NewUsersHandlers(srvc service.UserService) *UsersHandlers {
	return &UsersHandlers{srvc: srvc}
}

func (uh *UsersHandlers) GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	users, err := uh.srvc.GetAll(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		output, _ := json.Marshal(map[string]interface{}{
			"error": err,
		})

		w.Write(output)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	output, _ := json.Marshal(users)
	w.Write(output)
}

func (uh *UsersHandlers) GetById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	ctx := context.WithValue(r.Context(), "id", ps.ByName("id"))

	user, err := uh.srvc.GetById(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		output, _ := json.Marshal(map[string]interface{}{
			"error": err,
		})

		w.Write(output)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	output, _ := json.Marshal(user)
	w.Write(output)
}

package server

import (
	"context"
	"encoding/json"
	"github.com/open-ticket/user-service/models/user"
	"net/http"
)

type getAllUsersResponse struct {
	Users []user.User `json:"users"`
	Err   string      `json:"error,omitempty"`
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

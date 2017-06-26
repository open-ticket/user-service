package server

import (
	"context"
	"encoding/json"
	"net/http"
)

type getAllUsersRequest struct {
}

func decodeGetAllUsersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getAllUsersRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

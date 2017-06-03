package server

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

func makeGetAllUsersEndpoint(svc userService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		users, err := svc.getAllUsers()
		if err != nil {
			return getAllUsersResponse{nil, err.Error()}, nil
		}
		return getAllUsersResponse{users, ""}, nil
	}
}

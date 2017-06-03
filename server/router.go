package server

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
)

type urlEndpoint struct {
	path    string
	handler http.Handler
}

func setupEndpoints(svc userService) []urlEndpoint {
	var handlers []urlEndpoint
	getAllUsersHandler := httptransport.NewServer(
		makeGetAllUsersEndpoint(svc),
		decodeGetAllUsersRequest,
		encodeResponse,
	)

	handlers = append(handlers, urlEndpoint{"/user", getAllUsersHandler})

	return handlers
}

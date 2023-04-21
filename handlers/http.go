package handlers

import "go-final-project-mygram/services"

type HttpServer struct {
	service services.ServicesInterface
}

func NewHttpServer(service services.ServicesInterface) HttpServer {
	return HttpServer{service: service}
}

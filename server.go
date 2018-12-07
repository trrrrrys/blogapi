package main

import (
	"os"
	"blog-api/interface/api/rest"
	"blog-api/interface/api/middleware"
	"blog-api/application/usecase"
	"blog-api/infrastructure/datastore"
	"blog-api/interface/api/graphql/handler"
	"blog-api/interface/api/router"

	"log"
	"net/http"
)

func main() {
	cr := datastore.NewContentRepository(os.Getenv("PROJECT_ID"))
	ur := datastore.NewUserRepository()
	cu := usecase.NewContentUsecase(cr)
	uu := usecase.NewUserUsecase(ur)
	ch := handler.NewContentHandler(cu)
	uh := handler.NewUserHandler(uu)
	h := router.Route(handler.NewHandler(uh, ch), rest.NewContentHandler(cu))
	h = middleware.SetMiddleware(h)
	
	log.Println(http.ListenAndServe(":8080", h))
}

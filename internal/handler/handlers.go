package handler

import (
	"fmt"
	"forum/internal/service"
	"html/template"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		services: s,
	}
}

func (h *Handler) InitRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/auth/sign-up", h.signUp)
	mux.HandleFunc("/auth/sign-in", h.signIn)
	mux.HandleFunc("/log-out", h.logOut)
	mux.HandleFunc("/post/create", h.createPost)
	mux.HandleFunc("/my-posts", h.myPosts)
	return mux
}

func (h *Handler) templExecute(w http.ResponseWriter, path string, data interface{}) error {
	templ, err := template.ParseFiles(path)
	if err != nil {
		fmt.Println("ParseFiles()")
		h.ErrorPage(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return err
	}
	if err = templ.Execute(w, data); err != nil {
		fmt.Println("templExecute()")
		h.ErrorPage(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return err
	}
	return nil
}

package handler

import (
	"forum/internal/service"
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

package app

import (
	"log"
)

func Run(cfgFilePath string) error {
	cfg := config.NewConfig(cfgFilePath)
	// server := server.NewServer(cfg, handler)
	log.Printf("Starting server at port %v\nhttp://localhost:%v\n", cfg.Http.Addr, cfg.Http.Addr)
	return server.Srv.ListenAndServe()
}

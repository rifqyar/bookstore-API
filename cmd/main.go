package main

import (
	"log"

	"bookstrore-api/internal/config"
	"bookstrore-api/internal/db"
	"bookstrore-api/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load() // baca .env / env vars
	conn, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("db connect: %v", err)
	}
	defer db.Close(conn)

	r := gin.Default()
	routes.Register(r, conn, cfg)

	addr := ":" + cfg.AppPort
	log.Printf("listening on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
	"web-chat/api"
	"web-chat/config"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {

	if err := setupEnviroment(); err != nil {
		return err
	}
	log.Println("[server][setupEnviroment][done]")

	router, err := api.InitRouter()
	if err != nil {
		return err
	}
	log.Println("[server][InitRouter][done]")

	server := http.Server{
		Addr:           fmt.Sprintf(":%s", config.ServerConfig.Port),
		Handler:        router,
		ReadTimeout:    config.ServerConfig.ReadTimeout * time.Second,
		WriteTimeout:   config.ServerConfig.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err = server.ListenAndServe(); err != nil {
		return err
	}

	return nil

}

func setupEnviroment() error {
	env := flag.String("env", "dev", "To set environment dev/prod")

	flag.Parse()

	if *env != "dev" && *env != "prod" {
		return fmt.Errorf("invalid environment type. check --help for to check env options")
	}

	if *env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	if err := config.Setup(*env); err != nil {
		return err
	}
	log.Println("[server][config.Setup][done]", *env)

	return nil
}

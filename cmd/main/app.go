package main

import (
	"github.com/PavelDonchenko/40projects/rest-api-tutorial/internal/user"
	"github.com/PavelDonchenko/40projects/rest-api-tutorial/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("Create Router")
	router := httprouter.New()

	logger.Info("Register new handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	run(router)

}

func run(router *httprouter.Router) {
	logger := logging.GetLogger()
	logger.Info("Start application")
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	logger.Info("Server is listening port 1234")
	logger.Fatal(server.Serve(listener))
}

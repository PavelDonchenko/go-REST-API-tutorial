package main

import (
	"fmt"
	"github.com/PavelDonchenko/40projects/rest-api-tutorial/internal/config"
	"github.com/PavelDonchenko/40projects/rest-api-tutorial/internal/user"
	"github.com/PavelDonchenko/40projects/rest-api-tutorial/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("Create Router")
	router := httprouter.New()

	cfg := config.GetConfig()

	//cfgMongo := cfg.MongoDB
	//
	//mongoDbClient, err := mongodb.NewClient(context.Background(), cfgMongo.Host, cfgMongo.Port, cfgMongo.Username, cfgMongo.Password, cfgMongo.Database, cfgMongo.AuthDB)
	//if err != nil {
	//	panic(err)
	//}
	//storage := db.NewStorage(mongoDbClient, cfg.MongoDB.Collection, logger)

	logger.Info("Register new handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	run(router, cfg)

}

func run(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()
	logger.Info("Start application")

	var listener net.Listener
	var listenerError error

	if cfg.Listen.Type == "sock" {
		logger.Info("detected app path")
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info("Create socket")
		socketPath := path.Join(appDir, "app.sock")
		logger.Debugf("socket path:%s", socketPath)

		logger.Info("listen unix socket")
		listener, listenerError = net.Listen("unix", socketPath)
		logger.Infof("server is listening unix socket %s", socketPath)
	} else {
		logger.Info("listen  tcp")
		listener, listenerError = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIp, cfg.Listen.Port))
		logger.Infof("server is listening %s:%s", cfg.Listen.BindIp, cfg.Listen.Port)
	}
	if listenerError != nil {
		panic(listenerError)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	logger.Fatal(server.Serve(listener))
}

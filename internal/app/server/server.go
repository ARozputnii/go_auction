package server

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"go_auction/internal/pkg/config"
	"go_auction/internal/pkg/handlers"
	"go_auction/internal/pkg/models"
	"go_auction/internal/pkg/services"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
	Server *http.Server
}

func (s *App) Start() {
	config.InitEnvs(".env.dev")

	s.DB = config.InitDB()
	
	config.Migrate(s.DB)

	model := models.NewApplicationModel(s.DB)
	service := services.NewApplicationService(model)
	handler := handlers.NewApplicationHandler(service)

	s.Router = handler.InitRoutes()

	s.Server = s.initServer(s.Router)

	go func() {
		log.Fatal(s.Server.ListenAndServe())
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Info("TodoApp Shutting Down")

	if err := s.Server.Shutdown(context.Background()); err != nil {
		logrus.Fatalf("error occured on server shutting down: %s", err.Error())
	}

	sqlDB, _ := s.DB.DB()
	if err := sqlDB.Close(); err != nil {
		logrus.Fatalf("error occured on db connection close: %s", err.Error())
	}
}

func (s *App) initServer(router *mux.Router) *http.Server {
	port := os.Getenv("PORT")
	if port == "" {
		logrus.Fatalf("\nEnv 'PORT' is missing!")
	}

	adrr := fmt.Sprintf("127.0.0.1:%s", port)
	srv := &http.Server{
		Handler:      router,
		Addr:         adrr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logrus.Infof("Server is running on Port=%s\n", port)

	return srv
}

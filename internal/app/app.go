package app

import (
	"database/sql"
	"fmt"
	grpcapp "github.com/EvgeniyMdr/commentService/internal/app/grpc"
	"github.com/EvgeniyMdr/commentService/internal/config"
	"github.com/EvgeniyMdr/commentService/internal/db"
	"github.com/EvgeniyMdr/commentService/internal/repositories"
	"github.com/EvgeniyMdr/commentService/internal/services"
	"log"
)

type App struct {
	GRPCSrv *grpcapp.App
	Db      *sql.DB
}

func New() *App {
	mainConfig := config.NewServiceConfig()
	database, err := db.ConnectToDB(mainConfig.GetDbSettings())

	log.Printf("DB Ping successful. DB Stats: %+v\n", database.Stats())

	if err != nil {
		_ = fmt.Errorf("error: %w", err)
	}

	commentsRepo := repositories.NewCommentsRepository(database)

	commentsService := services.NewCommentsService(commentsRepo)

	grpcApp := grpcapp.New(commentsService)

	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	return &App{GRPCSrv: grpcApp, Db: database}
}

// TODO: Выяснить правильно ли я делаю стоп
func (a App) Stop() {
	defer func() {
		log.Printf("DB Close deffer successful.")
		if err := a.Db.Close(); err != nil {
			log.Fatalf("Ошибка закрытия соединения с базой данных: %v", err)
		}
	}()
}

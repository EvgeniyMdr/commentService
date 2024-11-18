package app

import (
	"fmt"
	grpcapp "github.com/EvgeniyMdr/commentService/internal/app/grpc"
	"github.com/EvgeniyMdr/commentService/internal/config"
	"github.com/EvgeniyMdr/commentService/internal/db"
	"github.com/EvgeniyMdr/commentService/internal/repositories"
	"github.com/EvgeniyMdr/commentService/internal/services"
	"log"
	"os"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New() *App {
	mainConfig := config.NewServiceConfig()
	database, err := db.ConnectToDB(mainConfig.GetDbSettings())
	migrationService := db.NewMigrationService(database)

	if err != nil {
		_ = fmt.Errorf("error: %w", err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Ошибка получения рабочей директории: %v", err)
	}
	log.Println("Текущая рабочая директория:", cwd)

	err = migrationService.Up("/app/db/migrations")

	if err != nil {
		log.Fatalf("Ошибка миграции в базе данных: %v", err)
	}

	commentsRepo := repositories.NewCommentsRepository(database)

	commentsService := services.NewCommentsService(commentsRepo)

	grpcApp := grpcapp.New(commentsService)

	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	defer func() {
		if err := database.Close(); err != nil {
			log.Fatalf("Ошибка закрытия соединения с базой данных: %v", err)
		}
	}()

	return &App{GRPCSrv: grpcApp}
}

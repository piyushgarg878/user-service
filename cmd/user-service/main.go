package main

import (
    "log"
    "user-service/internal/config"
    "user-service/internal/domain"
    "user-service/internal/handler"
    "user-service/internal/repository"
    "user-service/internal/router"
    "user-service/internal/service"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func main() {
	//setting up configurations
    cfg := config.LoadConfig()

	//setting up database
    db, err := gorm.Open(postgres.Open(cfg.DBUrl), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect DB:", err)
    }
    // Run migrations
    db.AutoMigrate(&domain.User{}, &domain.UserProfile{})

    // Setup repos
    userRepo := repository.NewUserRepository(db)
    profileRepo := repository.NewProfileRepository(db)

    // Service
    userService := service.NewUserService(userRepo, profileRepo)

    // Handler
    userHandler := handler.NewUserHandler(userService)

    // Router
    r := router.SetupRouter(userHandler)
    r.Run(":" + cfg.Port)
}
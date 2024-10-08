package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/oogway93/golangArchitecture/internal/adapter/config"
	"github.com/oogway93/golangArchitecture/internal/adapter/logger"
	repositoryPostgres "github.com/oogway93/golangArchitecture/internal/core/repository/postgres"
	"github.com/oogway93/golangArchitecture/internal/core/repository/postgres/auth"
	"github.com/oogway93/golangArchitecture/internal/core/repository/postgres/models"
	repositoryPostgresShop "github.com/oogway93/golangArchitecture/internal/core/repository/postgres/shop"
	repositoryPostgresUser "github.com/oogway93/golangArchitecture/internal/core/repository/postgres/user"
	repositoryRedis "github.com/oogway93/golangArchitecture/internal/core/repository/redis"

	HTTP "github.com/oogway93/golangArchitecture/internal/core/server/serverHTTP"

	serviceAuth "github.com/oogway93/golangArchitecture/internal/core/service/auth"
	serviceShop "github.com/oogway93/golangArchitecture/internal/core/service/shop"
	serviceUser "github.com/oogway93/golangArchitecture/internal/core/service/user"
)

func main() {
	gin.SetMode(gin.DebugMode)
	config, err := config.New()
	if err != nil {
		slog.Error("Error loading environment variables", "error", err)
		os.Exit(1)
	}

	// Set logger
	logger.Set(config.App)

	slog.Info("Starting the application", "app", config.App.Name, "env", config.App.Env)

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file",
			err)
	}

	db := repositoryPostgres.DatabaseConnection(repositoryPostgres.Config{
		Username: config.DB.User,
		Password: config.DB.Password,
		Host:     config.DB.Host,
		Port:     config.DB.Port,
		DBName:   config.DB.Name,
		SSLMode:  config.DB.SSLMode,
	})
	db.AutoMigrate(&models.User{}, &models.Category{}, &models.Product{}, &models.Order{}, &models.Delivery{}, &models.OrderItem{})
	slog.Info("Successfully migrated the database")

	addr := fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port)
	cache, err := repositoryRedis.New(repositoryRedis.Config{
		Addr:     addr,
		Password: config.Redis.Password,
		Expiration: time.Duration(config.Redis.Expiration) * time.Minute,
	})
	if err != nil {
		slog.Error("Error initializing cache connection", "error", err)
	}
	defer cache.Close()

	slog.Info("Successfully connected to the cache server")

	// Category
	categoryRepo := repositoryPostgresShop.NewRepositoryCategoryShop(db)
	categoryService := serviceShop.NewServiceShopCategory(categoryRepo, cache)

	// Product
	productRepo := repositoryPostgresShop.NewRepositoryProductShop(db)
	productService := serviceShop.NewServiceShopProduct(productRepo, cache)

	// Order
	orderRepo := repositoryPostgresShop.NewRepositoryOrderShop(db)
	orderService := serviceShop.NewServiceShopOrder(orderRepo, cache)

	// User
	userRepo := repositoryPostgresUser.NewRepositoryUser(db)
	userService := serviceUser.NewServiceUser(userRepo, cache)

	// Auth
	authRepo := repositoryPostgresAuth.NewRepositoryAuth(db)
	authService := serviceAuth.NewServiceAuth(authRepo, cache)

	router := HTTP.SetupRouter(categoryService, productService, orderService, userService, authService)

	server := new(HTTP.Server)
	if err := server.Run(config.HTTP.Port, router); err != nil {
		slog.Error("Some errors in initialization routes", "error", err)
	}
}

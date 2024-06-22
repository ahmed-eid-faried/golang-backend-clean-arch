package main

import (
	"log"
	"os"
	"time"

	"github.com/quangdangfit/gocommon/logger"
	"github.com/quangdangfit/gocommon/validation"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	// "golang.org/x/oauth2/google"

	// orderModel "main/internal/order/model"
	// productModel "main/internal/product/model"
	addressModel "main/internal/address/model"
	grpcServer "main/internal/server/grpc"
	httpServer "main/internal/server/http"
	userModel "main/internal/user/model"
	"main/pkg/config"
	"main/pkg/dbs"
	"main/pkg/redis"
)

//	@title			main Swagger API
//	@version		1.0
//	@description	Swagger API for main.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Quang Dang
//	@contact.email	quangdangfit@gmail.com

//	@license.name	MIT
//	@license.url	https://github.com/MartinHeinz/go-project-blueprint/blob/master/LICENSE

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

// @title User API
// @description API for user management
// @version 1.0
// @host localhost:8888
// @BasePath /api/v1
func main() {
	cfg := config.LoadConfig()
	logger.Initialize(cfg.Environment)
	//*********************************************

	// db, err := dbs.NewDatabase(cfg.DatabaseURI)
	// if err != nil {
	// 	logger.Fatal("Cannot connect to database", err)
	// }

	// *********************************************
	var db *dbs.Database
	var err error
	for i := 0; i < 10; i++ { // retry up to 10 times
		db, err = dbs.NewDatabase(cfg.DatabaseURI)
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database. Retrying in 5 seconds... (%d/10)\n", i+1)
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		logger.Fatal("Cannot connect to database", err)
		os.Exit(1)
	}
	//*********************************************
	oauthConfig := &oauth2.Config{
		ClientID:     cfg.GOOGLE_CLIENT_ID,
		ClientSecret: cfg.GOOGLE_CLIENT_SECRET,
		RedirectURL:  "http://localhost:8888/auth/google/callback",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}

	//*********************************************

	err = db.AutoMigrate(&userModel.User{}, &addressModel.Address{})
	if err != nil {
		logger.Fatal("Database migration fail", err)
	}

	validator := validation.New()

	cache := redis.New(redis.Config{
		Address:  cfg.RedisURI,
		Password: cfg.RedisPassword,
		Database: cfg.RedisDB,
	})

	go func() {
		httpSvr := httpServer.NewServer(validator, db, cache, oauthConfig)
		if err = httpSvr.Run(); err != nil {
			logger.Fatal(err)
		}
	}()

	grpcSvr := grpcServer.NewServer(validator, db, cache, oauthConfig)
	if err = grpcSvr.Run(); err != nil {
		logger.Fatal(err)
	}

}

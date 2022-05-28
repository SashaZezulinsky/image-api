package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/image-api/config"
	"github.com/image-api/internal/server"
	"github.com/image-api/pkg/utils"
)

func main() {
	rand.Seed(time.Now().Unix())

	log.Println("Starting api server")

	configPath := utils.GetConfigPath(os.Getenv("config"))

	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.MongoDB.MongoURI))
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	defer func() {
		if err := mongoClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	s := server.NewServer(cfg, mongoClient)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"basketsimulation/pkg/infrastructure/mongo"
	"basketsimulation/pkg/service"
	"basketsimulation/pkg/ui"
	"log"
	"os"
)

func main() {
	mongoClient := mongo.NewClient("mongodb://root:example@localhost:27017")
	matchRepository := mongo.NewMatchRepository(mongoClient, "basket-simulation")
	matchPlayerRepository := mongo.NewMatchPlayerRepository(mongoClient, "basket-simulation")
	go func() {
		service.NewMatchService(matchRepository, matchPlayerRepository).Start()
	}()
	server := ui.NewServer()
	if err := server.Run(matchRepository, matchPlayerRepository); err != nil {
		log.Print(err)
		os.Exit(1)
	}
}
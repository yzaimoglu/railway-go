package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/k0kubun/pp/v3"
	"github.com/yzaimoglu/railway-go/types"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Could not read environment variables from .env file")
	}
	token := os.Getenv("RAILWAY_TOKEN")
	railway := types.NewRailwayClient(token)
	res, err := railway.GetRegions()
	if err != nil {
		log.Fatalf("Could not get regions: %s", err.Error())
	}
	pp.Println(res)
}

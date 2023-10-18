package main

import (
	"os"

	"github.com/inocentini/codepix-go/application/grpc"
	"github.com/inocentini/codepix-go/infrastrutucte/db"
	"github.com/jinzhu/gorm"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartgRPCServer(database, 50051)
}

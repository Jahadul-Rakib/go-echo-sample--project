package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"sync"
)

type manager struct {
	Ctx      context.Context
	Database *mongo.Database
}

var dbInstance *manager
var onceManagerRun sync.Once
var DB_URL = os.Getenv("DB_URL")

func GetDmManager() *manager {
	onceManagerRun.Do(func() {
		log.Println("[INFO] Starting Initializing Singleton DB Manager")
		dbInstance = &manager{}
		dbInstance.initConnection()
	})
	return dbInstance
}

func (dm *manager) initConnection() {
	dm.Ctx = context.Background()
	clientOptions := options.Client().ApplyURI("mongodb://root:root@localhost:27017/")
	client, err := mongo.Connect(dm.Ctx, clientOptions)
	if err != nil {
		log.Println("Db Error: ", err.Error())
	} else {
		log.Println("Connected to MongoDB....!")
		dbInstance := client.Database("office-employee")
		dm.Database = dbInstance
	}
}

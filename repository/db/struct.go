package db

import (
	"github.com/AditWiBu69/pakage_be/config"
	mongo "github.com/kamagasaki/go-utils/mongo"
	"os"
)

var MongoString = os.Getenv("MONGOSTRCONNECT")

var DBATS = mongo.DBInfo{
	DBString: MongoString,
	DBName:   config.ATSMDB,
}

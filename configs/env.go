package configs 

import (
	"github.com/joho/godotenv" 
	"os" 
	"log"
)

func EnvMongoURI() string {
	err:= godotenv.Load() 

	if err != nil {
		log.Fatal("error in .env file => ", err)
	}

	return os.Getenv("MongoUri")
}
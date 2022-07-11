package configs 

import (
	"github.com/joho/godotenv" 
	"os" 
	"log"
)

func EnvMongoURI() string {
	error:= godotenv.Load() 

	if error != nil {
		log.Fatal("error in .env file => ", error)
	}

	return os.Getenv("MongoUri")
}
